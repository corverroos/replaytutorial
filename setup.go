package replaytutorial

import (
	"context"
	"database/sql"
	"flag"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/corverroos/replay"
	replay_server "github.com/corverroos/replay/server"
	"github.com/corverroos/truss"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"

	"github.com/corverroos/replaytutorial/lib/filenotifier"
)

var dbURL = flag.String("db_url", "mysql://root@unix(/tmp/mysql.sock)/", "mysql db url (without DB name)")
var dbName = flag.String("db_name", "replay_tutorial", "db schema name")
var dbRestart = flag.Bool("db_refresh", false, "Whether to drop the existing DB on startup")
var runLoops = flag.Bool("server_loops", true, "Whether to not to run the replay server loops. Note only a single active process may run the server loops")

type State struct {
	DBC        *sql.DB
	Replay     replay.Client
	Cursors    reflex.CursorStore
	AppCtxFunc func() context.Context
}

// Main calls the mainFunc with the tutorial state. It support additional application logic sql migrations.
func Main(mainFunc func(context.Context, State) error, migrations ...string) {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	dbc, err := connectDB(ctx, migrations...)
	if err != nil {
		log.Error(ctx, err)
		os.Exit(1)
	}

	cstore := rsql.NewCursorsTable("cursors", rsql.WithCursorAsyncDisabled()).ToStore(dbc)

	notifier, err := filenotifier.New()
	if err != nil {
		log.Error(ctx, errors.Wrap(err, "notifier"))
		os.Exit(1)
	}

	rcl := replay_server.NewDBClient(dbc, rsql.WithEventsNotifier(notifier))

	appCtxFunc := func() context.Context {
		return ctx
	}

	if *runLoops {
		rcl.StartLoops(appCtxFunc, cstore, "")
	}

	rand.Seed(time.Now().UnixNano())

	state := State{
		DBC:        dbc,
		Replay:     rcl,
		Cursors:    cstore,
		AppCtxFunc: appCtxFunc,
	}

	err = mainFunc(ctx, state)
	if ctx.Err() != nil {
		log.Info(ctx, "app context canceled")
	} else if err != nil {
		log.Error(ctx, err)
		os.Exit(1)
	}
}

// AwaitComplete streams all events and returns on the first RunCompleted event.
func AwaitComplete(ctx context.Context, cl replay.Client, run string) error {
	sc, err := cl.Stream("", "", "")(ctx, "")
	if err != nil {
		return err
	}

	for {
		e, err := sc.Recv()
		if err != nil {
			return err
		}

		// Use the replay event handling functions.
		err = replay.Handle(e, replay.HandleRunCompleted(func(_, _, r string) error {
			if run == r {
				return io.EOF
			}
			return nil
		}))
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func connectDB(ctx context.Context, migrations ...string) (*sql.DB, error) {
	dbc, err := truss.Connect(*dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "initial db connect")
	}

	if *dbRestart {
		_, err := dbc.ExecContext(ctx, "DROP DATABASE IF EXISTS "+*dbName)
		if err != nil {
			return nil, err
		}
	}

	_, err = dbc.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+*dbName+" CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err != nil {
		return nil, err
	}

	dbc, err = truss.Connect(*dbURL + *dbName + "?parseTime=true")
	if err != nil {
		return nil, errors.Wrap(err, "db connect")
	}

	err = truss.Migrate(ctx, dbc, append(defaultMigrations, migrations...))
	if err != nil {
		return nil, errors.Wrap(err, "migrate")
	}

	return dbc, nil
}

var defaultMigrations = []string{`
CREATE TABLE replay_events (
  id BIGINT NOT NULL AUTO_INCREMENT,
  ` + "`key`" + ` VARCHAR(512) NOT NULL,
  namespace VARCHAR(255) NOT NULL, 
  workflow VARCHAR(255) NOT NULL,
  run VARCHAR(255),
  iteration INT NOT NULL,
  type INT NOT NULL,
  timestamp DATETIME(3) NOT NULL,
  message MEDIUMBLOB,

  PRIMARY KEY (id),
  UNIQUE by_type_key (type, ` + "`key`" + `),
  INDEX (type, namespace, workflow, run, iteration)
);
`, `
CREATE TABLE replay_sleeps (
  id BIGINT NOT NULL AUTO_INCREMENT,
  ` + "`key`" + ` VARCHAR(512) NOT NULL,
  created_at DATETIME(3) NOT NULL,
  complete_at DATETIME(3) NOT NULL,
  completed BOOL NOT NULL,

  PRIMARY KEY (id),
  UNIQUE by_key (` + "`key`" + `),
  INDEX (completed, complete_at)
);
`, `
CREATE TABLE replay_signal_awaits (
  id BIGINT NOT NULL AUTO_INCREMENT,
  ` + "`key`" + ` VARCHAR(512) NOT NULL,
  created_at DATETIME(3) NOT NULL,
  timeout_at DATETIME(3) NOT NULL,
  status TINYINT NOT NULL,

  PRIMARY KEY (id),
  UNIQUE by_key (` + "`key`" + `),
  INDEX (status)
);
`, `
CREATE TABLE replay_signals (
  id BIGINT NOT NULL AUTO_INCREMENT,
  namespace VARCHAR(255) NOT NULL,
  hash BINARY(128) NOT NULL,
  workflow VARCHAR(255) NOT NULL,
  run VARCHAR(255) NOT NULL,
  type TINYINT NOT NULL,
  external_id VARCHAR(255) NOT NULL,
  message MEDIUMBLOB,
  created_at DATETIME(3) NOT NULL,
  check_id BIGINT,

  PRIMARY KEY (id),
  UNIQUE uniq (namespace, hash)
);
`, `
CREATE TABLE cursors (
   id VARCHAR(255) NOT NULL,
   last_event_id BIGINT NOT NULL,
   updated_at DATETIME(3) NOT NULL,

   PRIMARY KEY (id)
);
`,
}

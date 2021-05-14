package replaytutorial

import (
	"context"
	"testing"
	"time"

	"github.com/luno/jettison/jtest"
	"github.com/stretchr/testify/require"
)

func TestTimestamps(t *testing.T) {
	*dbRestart = true
	*dbName = "tut_test"

	Main(func(ctx context.Context, state State) error {
		dbc := state.DBC

		_, err := dbc.ExecContext(ctx, "create table testtime (id bigint auto_increment, ts datetime(3) not null, primary key (id))")
		jtest.RequireNil(t, err)

		t0 := time.Now()
		_, err = dbc.ExecContext(ctx, "insert into testtime set ts=now(3)")
		jtest.RequireNil(t, err)

		var ts time.Time
		err = dbc.QueryRowContext(ctx, "select ts from testtime where id=1").Scan(&ts)
		jtest.RequireNil(t, err)
		require.InDelta(t, t0.UnixNano(), ts.UnixNano(), 1e9) // Within 1 sec

		t0 = time.Now()
		_, err = dbc.ExecContext(ctx, "insert into testtime set ts=?", t0)
		jtest.RequireNil(t, err)

		err = dbc.QueryRowContext(ctx, "select ts from testtime where id=2").Scan(&ts)
		jtest.RequireNil(t, err)
		require.InDelta(t, t0.UnixNano(), ts.UnixNano(), 1e9) // Within 1 sec

		return nil
	})

}

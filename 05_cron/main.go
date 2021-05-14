// Exercise 05_cron guides you through designing a cron-like workflow that triggers logic on each interval.
// Note that this is a long running workflow, so it uses the Restart feature.
package main

import (
	"context"
	"time"

	"github.com/corverroos/replay/typedreplay"
	"github.com/luno/fate"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"

	tut "github.com/corverroos/replaytutorial" // Alias replaytutorial to tut for brevity.
)

// Step 0: main functions always just call tut.Main(Main).
func main() {
	tut.Main(Main)
}

// Step 1: Replay always requires protobufs, so generate your types.
//go:generate protoc --go_out=plugins=grpc:. ./pb.proto

// Step 2: typedreplay requires a locally defined Backends.

type Backends struct{}

// Step 3: Define typedreplay namespace

var _ = typedreplay.Namespace{
	Name: "05_longhello",
	// Workflows define the workflow function names and types.
	Workflows: []typedreplay.Workflow{
		{
			Name:        "cron",
			Description: "Cron workflow calls the task activity on each interval",
			Input:       new(Timestamp),
		},
	},
	Activities: []typedreplay.Activity{
		{
			Name:        "task",
			Description: "This task is called for each interval",
			Func:        ExecTask,
		},
	},
}

// Step 4: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

func ExecTask(ctx context.Context, b Backends, f fate.Fate, ts *Timestamp) (*Empty, error) {
	log.Info(ctx, "exec task", j.KV("timestamp", time.Unix(ts.UnixSec, 0)))
	return new(Empty), nil
}

func Cron(flow cronFlow, ts *Timestamp) {
	now := flow.Now() // Note that due to deterministic requirements time.Now is not very accurate and uses LastEvent().Timestamp.
	thisRun := time.Unix(ts.UnixSec, 0)

	// Sleep until this run.
	if thisRun.After(now) {
		flow.Sleep(thisRun.Sub(now))
	}

	// Execute the task
	flow.ExecTask(ts)

	// Calculate next run. This is your cron schedule.
	nextRun := thisRun.Add(time.Second * 2)

	// Restart the run.
	flow.Restart(&Timestamp{UnixSec: nextRun.Unix()})
}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops, note it defines the signature of the Hello typed workflow function.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, Cron)

	ok, err := RunCron(ctx, s.Replay, "cron", &Timestamp{UnixSec: time.Now().Truncate(time.Second).Unix()})
	if err != nil {
		return err
	} else if ok {
		log.Info(ctx, "started cron for the first time")
	} else {
		log.Info(ctx, "cron already started")
	}

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 6: Run the program and confirm the same expected output.
//go:generate go run github.com/corverroos/replaytutorial/05_cron

// Example output:
//  I 15:31:39.635 05_cron/main.go:86: started cron for the first time
//  I 15:31:39.636 05_cron/main.go:91: Press Ctrl-C to exit...
//  I 15:31:39.637 05_cron/main.go:53: exec task[consumer=replay_activity/04_cronactivity/task,replay_run=cron,timestamp=2021-05-14 15:31:39 +0200 SAST]
//  I 15:31:41.645 05_cron/main.go:53: exec task[consumer=replay_activity/04_cronactivity/task,replay_run=cron,timestamp=2021-05-14 15:31:41 +0200 SAST]
//  I 15:31:43.659 05_cron/main.go:53: exec task[consumer=replay_activity/04_cronactivity/task,replay_run=cron,timestamp=2021-05-14 15:31:43 +0200 SAST]

// Step 7: Experiments
// - What happens if the program is not running for a couple of iterations and then started again?
// - How would you change the task activity to skip stale iterations?
// - What would happen if the activity blocked for multiple iterations? Would the cron start lagging?
// - What would be the difference if the workflow emitted an output per interval instead of calling an activity.
// - Since RunCreated event contains the Timestamp message, same as the output, we could just skip the output and consume RunCreated events directly.

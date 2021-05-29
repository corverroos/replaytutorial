// Exercise 06_appendonlylog guides you through using replay as an append-only-log. Since
// it is possible to directly consume RunCreated events (with their data) from
// application logic via a reflex consumer, modelling an append-only-log is as simple as defining
// a worklfow with only an input and without any logic, or activities, or outputs.
package main

import (
	"context"

	"github.com/corverroos/replay/typedreplay"
	"github.com/luno/fate"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"

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
	Name: "06_appendonlylog",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "append",
			Description: "Appends data to the event log on creation (no actual logic)",
			Input:       new(Data),
		},
	},
	// No activities or outputs since application logic directly consumes the RunCreated events as the log.
}

// Step 4: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// noop workflow function.
func noop(flow appendFlow, ts *Data) {}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops.
	// Note technically one doesn't even need to start the workflow consumer loop since it doesn't do anything.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, noop)

	// TODO(you): Append some events to the log by calling RunAppend. Note that unique run IDs are required.

	// Define the log consume function
	_ = func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// TODO(you): Use the generated HandleAppendRun to consume the log entries
		// and just print the values
		panic("implement me")
	}

	// TODO(you): Define and run the reflex spec using the generated StreamAppend function that streams entries of append-only-log.

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 6: Run the program and confirm the same expected output.
//go:generate go run github.com/corverroos/replaytutorial/06_appendonlylog

// Step 7: Experiments
// - Add multiple independent log entry consumers

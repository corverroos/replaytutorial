// Example 00_helloworld shows how the replaytutorial package is used in a simple golang application to
// write a hello world implementation using the replay framework.
//
// Note: We are using the type-unsafe replay API in this first example, we will use the preferred
// typedreplay code generator later.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/corverroos/replay"
	"github.com/google/uuid"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
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

// Step 2: Define the workflow and activity functions

// Hello is a workflow that calls the Print with "Hello {name}" after sleeping 2 seconds.
func Hello(rc replay.RunContext, name *String) {
	rc.Sleep(2 * time.Second)

	msg := fmt.Sprintf("Hello %s", name.Value)

	rc.ExecActivity(Print, &String{Value: msg})
}

// Print is an activity function that logs the message value
func Print(ctx context.Context, state tut.State, f fate.Fate, message *String) (*Empty, error) {
	log.Info(ctx, message.Value)
	return new(Empty), nil
}

// Step 3: Define your Main function.
func Main(ctx context.Context, s tut.State) error {
	// All replay workflows and activities are grouped by namespace, so define it at the top.
	const namespace = "00_helloworld"

	// Register the activity and workflow consumers that will do the async processing.
	replay.RegisterWorkflow(s.AppCtxFunc, s.Replay, s.Cursors, namespace, Hello)
	replay.RegisterActivity(s.AppCtxFunc, s.Replay, s.Cursors, s, namespace, Print)

	// Run the workflow with a specific value. Please note:
	// - The workflow name is inferred from the above registered function name.
	// - It is up to the user to define the run identifier, in this case we use a random UUID.
	// - The proto message type must match that used by the workflow function.
	// - The typedreplay code generator covered later makes this much safer.
	run := uuid.New().String()

	ok, err := s.Replay.RunWorkflow(ctx, namespace, "Hello", run, &String{Value: "world"})
	if err != nil {
		return err
	} else if !ok {
		return errors.New("run already exists, duplicate UUID?!")
	}

	log.Info(ctx, "started run", j.KS("run", run))

	// Wait for the workflow run to complete. Note that it should still complete
	// even if the binary is restarted at this point.
	return tut.AwaitComplete(ctx, s.Replay, run)
}

// Step 4: Run the program and confirm the expected output
//go:generate go run github.com/corverroos/replaytutorial/00_helloworld

// Example output:
//  I 13:45:28.757 00_helloworld/main.go:71: started run[run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]
//  I 13:45:30.770 00_helloworld/main.go:44: Hello world[consumer=replay_activity/00_helloworld/Print,replay_run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]

// Note replay is robust to restarts, so early exit (Ctrl-C) should result in "Hello world" to be printed in the subsequent run.

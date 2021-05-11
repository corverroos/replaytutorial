// Exercise 02_signalhello guides you through using signals to send data to an active workflow run.
package main

import (
	"context"
	"fmt"

	"github.com/corverroos/replay/typedreplay"
	"github.com/google/uuid"
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

// Step 2: typedreplay requires a locally defined Backends. Since we do not actually use if now, it can be empty.

type Backends struct{}

// Step 3: Define typedreplay namespace

var _ = typedreplay.Namespace{
	// Name of the namespace is equivalent to the previously manually hardcoded 00_helloworld const.
	Name: "02_signalhello",
	// Workflows define the workflow function names and types.
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow just prints 'Hellow {signal}'",
			Input:       new(Empty), // Note that we do not provide the name at the start of the run, but later via a signal
			Signals: []typedreplay.Signal{
				{
					// Define the signal we want to send to the workflow.
					Name:        "name",
					Description: "Send the name to the workflow to print",
					Enum:        1,
					Message:     new(String),
				},
			},
		},
	},
	// Activities define the namespace activity function names and types.
	Activities: []typedreplay.Activity{
		// TODO(you): Define the Print activity.
	},
}

// Step 3: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// TODO(you): Define a similar Print activity function to 00_helloworld.
// func Print(...) {}

// TODO(you): Define a similar workflow function to 01_typedworld, except this time call AwaitName.
func Hello(flow helloFlow, _ *Empty) {
}

// Step 4: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops, note it defines the signature of the Hello typed workflow function.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, Hello)

	// Run the workflow with a specific value using the generated API.
	run := uuid.New().String()

	ok, err := RunHello(ctx, s.Replay, run, new(Empty))
	if err != nil {
		return err
	} else if !ok {
		return errors.New("Main already exists, duplicate UUID?!")
	}

	log.Info(ctx, "started run", j.KS("Main", run))

	// Ask the user to enter the name and then signal the run.
	var input string
	fmt.Print("Enter name:")
	if _, err = fmt.Scanln(&input); err != nil {
		return errors.Wrap(err, "read input")
	}

	// TODO(you): signal the run using the generated signal func SignalHelloName
	//  Note that external ID is a unique idempotency identifier for the signal. It can be anything for now.

	// Wait for the workflow Main to complete. Note that it should still complete
	// even if the binary is restarted at this point.
	return tut.AwaitComplete(ctx, s.Replay, run)
}

// Step 5: Run the program and confirm the same expected output as 00_helloworld
//go:generate bash -c "echo world | go run github.com/corverroos/replaytutorial/02_signalhello"

// Example output:
//  I 13:45:28.757 00_helloworld/main.go:71: started run[run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]
//  Enter name: world
//  I 13:45:30.770 00_helloworld/main.go:44: Hello world[consumer=replay_activity/00_helloworld/Print,replay_run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]

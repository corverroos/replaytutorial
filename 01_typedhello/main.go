// Exercise 01_typedhello guides you through using the typedreplay code generator to implement a
// type-safe version of 00_helloworld
package main

import (
	"context"

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

// Step 2: typedreplay requires a locally defined Backends dependency wrapper type.

type Backends struct {
	tut.State
}

// Step 3: Define typedreplay namespace

var _ = typedreplay.Namespace{
	// Name of the namespace is equivalent to the previously manually hardcoded 00_helloworld const.
	Name: "01_typedhello",
	// Workflows define the workflow function names and types.
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow just prints 'Hello {name}'",
			Input:       new(String),
		},
	},
	// Activities define the namespace activity function names and types.
	Activities: []typedreplay.Activity{
		// TODO(you): Define the Print activity.
	},
}

// Step 4: Install the typedreplay tool locally
//go:generate go install github.com/corverroos/replay/typedreplay/cmd/typedreplay@latest

// And then generate the type-safe replay API for the above definition.
//go:generate typedreplay

// TODO(you): Define a similar Print activity function to 00_helloworld.
// func Print(...) {}

// TODO(you): Define a similar workflow function to 00_helloworld, except this time, it is all type-safe.
func Hello(flow helloFlow, name *String) {
}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops, note it defines the signature of the Hello typed workflow function.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{s}, Hello)

	// Run the workflow with a specific value using the generated API.
	run := uuid.New().String()

	ok, err := RunHello(ctx, s.Replay, run, &String{Value: "world"})
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

// Step 6: Run the program and confirm the same expected output as 00_helloworld
//go:generate go run github.com/corverroos/replaytutorial/01_typedhello

// Example output:
//  I 13:45:28.757 00_helloworld/main.go:71: started run[run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]
//  I 13:45:30.770 00_helloworld/main.go:44: Hello world[consumer=replay_activity/00_helloworld/Print,replay_run=ba845343-58ba-4d88-8b8b-c46c8f66d25f]

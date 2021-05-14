// Exercise 03_outputhello guides you through emitting outputs in a pub/sub style from an active workflow run
// as well how to consume these outputs from application logic.
//
// From the replay docs:
//   Outputs vs Activities: Both activities and outputs can be used by workflows to trigger business
//   logic with data.
//   An activity's input, logic and output are tightly coupled with a workflow (think function calls).
//   While an output is only data emitted by a workflow decoupled from consuming logic (think pub/sub).
//   Another big benefit of outputs are that they do not impact workflow determinism in replay; that
//   means outputs may be added to, reordered in, or removed from, active runs. It is therefore recommended
//   to use outputs over activities where possible.
package main

import (
	"context"

	"github.com/corverroos/replay/typedreplay"
	"github.com/google/uuid"
	"github.com/luno/fate"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

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
	Name: "03_outputhello",
	// Workflows define the workflow function names and types.
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow outputs the message 'Hello {name}'",
			Input:       new(String),
			Outputs: []typedreplay.Output{
				{
					// Define the message output we want to emit from the workflow.
					Name:        "message",
					Description: "Message emitted from the workflow",
					Message:     new(String),
				},
			},
		},
	},
	// The message output is consumed by application logic, so we do not include any replay activities.
}

// Step 4: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// TODO(you): Define a similar workflow function to 01_typedworld, except this time call EmitMessage.
func Hello(flow helloFlow, name *String) {
}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops, note it defines the signature of the Hello typed workflow function.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, Hello)

	run := uuid.New().String()

	// TODO(you): Run the workflow with a specific value using the generated API.

	log.Info(ctx, "started run", j.KS("run", run))

	// Define a reflex consume function that will print message outputs.
	// Instead of a replay activity, outputs are processed in application logic.
	consume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Use the generated HandleMessage functional option.
		return HandleMessage(e, func(r string, message *String) error {
			log.Info(ctx, message.Value, j.KS("replay_run", r)) // Now we can just print here

			// Notify that we are done
			if r == run {
				log.Info(ctx, "Press Ctrl-C to exit...")
			}
			return nil
		})
	}

	// Define and run the reflex spec using the generated StreamHello function that streams events of the hello workflow.
	spec := reflex.NewSpec(
		StreamHello(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("03_outputhello/print", consume))

	go rpatterns.RunForever(s.AppCtxFunc, spec)

	<-ctx.Done()
	return nil
}

// Step 6: Run the program and confirm the same expected output as 00_helloworld
//go:generate go run github.com/corverroos/replaytutorial/03_outputhello

// Example output:
//  I 13:48:17.765 03_outputhello/main.go:88: started run[run=f32b1626-e5af-4f8e-a62c-796804c38074]
//  I 13:48:19.796 03_outputhello/main.go:95: Hello world[consumer=03_outputhello/print,replay_run=f32b1626-e5af-4f8e-a62c-796804c38074]
//  I 13:48:19.796 03_outputhello/main.go:99: Press Ctrl-C to exit...[consumer=03_outputhello/print]

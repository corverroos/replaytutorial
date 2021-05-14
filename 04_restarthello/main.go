// Exercise 04_restarthello guides you through use the Restart feature for writing repetitive long-lived workflows.
//
// Replay workflow runs are deterministic and event-driven; given the same input argument, the same internal state will be
// achieved since the same activities must be requested and the same activity results are returned. This allows replay
// to bootstrap an active run after a long sleep or after a restart.
//
// When bootstrapping an active run all previous events of that run are queried and replayed during bootstrapping.
// This obviously doesn't scale infinitely. Another reason to limit the amount of events of a run is to allow compaction
// or deletion of old events in the event log. As long as a run is active, the events cannot be deleted. Limiting the
// number of events related to run is therefore a important.
//
// Replay therefore supports "restarting" active runs. Restarting effectively exits the workflow and enters it again
// with the provided input argument. The run therefore restarts itself, clearing all internal state and starting
// fresh with an new input. Only the latest events after a restart is required for bootstrapping a restarted run. Events related
// to previous iterations of the run can be safely ignored or even deleted if required.
//
// In many cases, using restart to model an iterative/cyclic workflow also greatly simplifies the logic and makes it much
// simpler to reason about.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/corverroos/replay/typedreplay"
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
	Name: "04_repeathello",
	// Workflows define the workflow function names and types.
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow iteratively outputs the message 'Hello {name} #{iter}'",
			Input:       new(State),
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

// TODO(you): Define a similar workflow function to 01_typedworld, except this time call EmitMessage then Restart.
func Hello(flow helloFlow, state *State) {
	flow.Sleep(time.Second * 2)
	flow.EmitMessage(&String{Value: fmt.Sprintf("Hellow %s, #%d", state.Name, state.Iter)})
	state.Iter++
	flow.Restart(state)
}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops, note it defines the signature of the Hello typed workflow function.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, Hello)

	// TODO(you): Run the workflow. Note that you only need to ensure the same run has been started, since it is long lived.
	run := "singleton"

	ok, err := RunHello(ctx, s.Replay, run, &State{Name: "world"})
	if err != nil {
		return err
	} else if !ok {
		log.Info(ctx, "run already run")
	} else {
		log.Info(ctx, "started run")
	}

	// TODO(you): Define a reflex consume function that will print message outputs.
	consume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Use the generated HandleMessage functional option.
		_, err := HandleMessage(e, func(r string, message *String) error {
			log.Info(ctx, message.Value, j.KS("replay_run", r)) // Now we can just print here
			return nil
		})
		return err
	}

	// Define and run the reflex spec using the generated StreamHello function that streams events of the hello workflow.
	spec := reflex.NewSpec(
		StreamHello(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("04_restarthello/print", consume))

	go rpatterns.RunForever(s.AppCtxFunc, spec)

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 6: Run the program and confirm the same expected output
//go:generate go run github.com/corverroos/replaytutorial/04_restarthello

// Example output:
//  I 15:10:00.136 04_repeathello/main.go:96: started run
//  I 15:10:00.136 04_repeathello/main.go:117: Press Ctrl-C to exit...
//  I 15:10:02.147 04_repeathello/main.go:103: Hellow world, #0[consumer=04_repeathello/print,replay_run=singleton]
//  I 15:10:05.156 04_repeathello/main.go:103: Hellow world, #1[consumer=04_repeathello/print,replay_run=singleton]
//  I 15:10:08.168 04_repeathello/main.go:103: Hellow world, #2[consumer=04_repeathello/print,replay_run=singleton]

// Step 7: Do some experiments
//  - What happens if you change the run identifier? Can you get multiple active long-lived runs?
//  - Try adding the -db_refresh flag to start with empty DB.
//  - How could you exit a run after 10 iterations?
//  - Achieve the same result without emitting the message outputs. Instead consume RunCreated events
//    directly from the reflex consumer application logic via replay.Handle(e, replay.HandleRunCreated(...))
//  - Get a closer look at the replay event internals by handling all event types replay.Handle(e, ...)

// Exercise 09_parentjob guides you through using replay to model a parent job request
// that is claimed an executed by one of the child worker implementations.
//
// In system design a common problem is how model a job that is executed by one of
// multiple different implementations. Some examples include:
//  - Payment providers executing a payment
//  - Liquidity providers executing a order/trade.
//  - Cloud APIs creating a resource.
//
// A naive solution would be to define a common interface that the child workers (providers)
// should implement and then to execute the job directly from the parent worker using the interface.
// This is hard in practice since a common interface that is compatible (and works well) with
// all provider idiosyncrasies is usually not feasible and results in leaky/weak/brittle abstractions.
//
// The pattern proposed by this exercise uses a more declarative data driven design. The parent
// job request merely provides the input data as well as a simple state machine modelling
// the request life cycle. Child implementation listen for new requests, claims the request
// if applicable, executes the request in its own fashion, and then updates the parent state once complete.
// This inverses the dependency graph with child implementation knowning about the parent but
// not the other way round.
package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/corverroos/replay/typedreplay"
	"github.com/google/uuid"
	"github.com/luno/fate"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	tut "github.com/corverroos/replaytutorial" // Alias replaytutorial to tut for brevity.
)

// Increase showme to 1 to unhide the next solution.
//go:generate go run ../lib/showme 0

// Step 0: main functions always just call tut.Main(Main).

func main() {
	tut.Main(Main)
}

// Step 1: Replay always requires protobufs, so generate your types.
//go:generate protoc --go_out=plugins=grpc:. ./pb.proto

// Step 2: typedreplay requires a locally defined Backends.

type Backends struct{}

// Step 3: Define activity functions (there are no activities actually).

// Step 4: Define the typedreplay namespace

var _ = typedreplay.Namespace{
	Name: "09_parentjob",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "parent",
			Description: "Parent job request state machine",
			Input:       new(JobRequest),
			Signals: []typedreplay.Signal{
				{
					Name:        "claim_job",
					Description: "Claims the job request by a child provider.",
					Enum:        1,
					Message:     new(ClaimJob),
				}, {
					Name:        "complete_job",
					Description: "Completes the job request by a child provider.",
					Enum:        2,
					Message:     new(CompleteJob),
				},
			},
		},
	},
}

// Step 5: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// Step 6: Define the buffer workflow function

func parent(flow parentFlow, req *JobRequest) {
	// TODO(you): Implement the buffer workflow state machine (hint: it is just a single switch on req.Status).
	//showme:hidden workflow
}

// Step 7: Define your Main function which is equivalent to a main function, just with some prepared state.

func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, parent)

	run := uuid.New().String()
	isFoo := rand.Float64() < 0.5
	isBar := !isFoo
	isBoth := rand.Float64() < 0.3
	log.Info(ctx, "creating job request", j.MKV{"is_foo": isFoo, "is_bar": isBar, "is_both": isBoth, "run": run})

	ok, err := RunParent(ctx, s.Replay, run, &JobRequest{
		Status: JobStatus_PENDING,
		IsFoo:  isFoo || isBoth,
		IsBar:  isBar || isBoth,
		Value:  fmt.Sprint(rand.Int63()),
	})
	if err != nil {
		return err
	} else if !ok {
		log.Info(ctx, "run already started")
	} else {
		log.Info(ctx, "new run started")
	}

	// TODO(you): Define the RunCreated event consume functions for the foo and bar providers that
	//  claims jobs with IsFoo/IsBar=true and prints "executing" and completes if claim successful.
	fooConsume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		//showme:hidden foo consume
		panic("implement me")
	}

	barConsume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		//showme:hidden bar consume
		panic("implement me")
	}

	// Define and run the reflex spec using the generated StreamMakerOrder function.
	fooSpec := reflex.NewSpec(
		StreamParent(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("09_parentjob/foo", fooConsume))

	barSpec := reflex.NewSpec(
		StreamParent(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("09_parentjob/bar", barConsume))

	go rpatterns.RunForever(s.AppCtxFunc, fooSpec)
	go rpatterns.RunForever(s.AppCtxFunc, barSpec)

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 8: Run the program and confirm the same expected output.
//go:generate go run github.com/corverroos/replaytutorial/09_parentjob

// Example output:
// I 16:27:18.927 09_parentjob/main.go:137: creating job request[is_bar=true,is_both=true,is_foo=false,run=ae3a3dc0-d0a1-4cfe-bf3f-a725b06da464]
// I 16:27:18.929 09_parentjob/main.go:150: new run started
// I 16:27:18.929 09_parentjob/main.go:208: Press Ctrl-C to exit...
// I 16:27:18.929 09_parentjob/main.go:180: bar claiming job[consumer=09_parentjob/bar,run=ae3a3dc0-d0a1-4cfe-bf3f-a725b06da464]
// I 16:27:18.930 09_parentjob/main.go:159: foo claiming job[consumer=09_parentjob/foo,run=ae3a3dc0-d0a1-4cfe-bf3f-a725b06da464]
// I 16:27:19.949 09_parentjob/main.go:163: foo executing job!![consumer=09_parentjob/foo,run=ae3a3dc0-d0a1-4cfe-bf3f-a725b06da464]
// I 16:27:20.962 09_parentjob/main.go:167: job complete[consumer=09_parentjob/foo,run=ae3a3dc0-d0a1-4cfe-bf3f-a725b06da464]

// Step 9: Experiments
// - Each provider can also be a its own workflow, what would that look like?

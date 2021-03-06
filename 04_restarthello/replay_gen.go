package main

// Code generated by typedreplay. DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/corverroos/replay"
	"github.com/golang/protobuf/proto"
	"github.com/luno/reflex"
	// TODO(corver): Support importing other packages.
)

const (
	_ns            = "04_repeathello"
	_wHello        = "hello"
	_oHelloMessage = "message"
)

// RunHello provides a type API for running the hello workflow.
// It returns true on success or false on duplicate calls or an error.
func RunHello(ctx context.Context, cl replay.Client, run string, message *State) (bool, error) {
	return cl.RunWorkflow(ctx, _ns, _wHello, run, message)
}

// startReplayLoops registers the workflow and activities for typed workflow functions.
func startReplayLoops(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore, b Backends,
	hello func(helloFlow, *State)) {

	helloFunc := func(ctx replay.RunContext, message *State) {
		hello(helloFlowImpl{ctx}, message)
	}
	replay.RegisterWorkflow(getCtx, cl, cstore, _ns, helloFunc, replay.WithName(_wHello))

}

// helloFlow defines a typed API for the hello workflow.
type helloFlow interface {

	// Sleep blocks for at least d duration.
	// Note that replay sleeps aren't very accurate and
	// a few seconds is the practical minimum.
	Sleep(d time.Duration)

	// CreateEvent returns the reflex event that started the run iteration (type is internal.CreateRun).
	// The event timestamp could be used to reason about run age.
	CreateEvent() *reflex.Event

	// LastEvent returns the latest reflex event (type is either internal.CreateRun or internal.ActivityResponse).
	// The event timestamp could be used to reason about run age.
	LastEvent() *reflex.Event

	// Run returns the run name/identifier.
	Run() string

	// Restart completes the current run iteration and starts a new run iteration with the provided input message.
	// The run state is effectively reset. This is handy to mitigate bootstrap load for long running tasks.
	// It also allows updating the activity logic/ordering.
	Restart(message *State)

	// EmitMessage stores the message output in the event log and returns when successful.
	EmitMessage(message *String)
}

type helloFlowImpl struct {
	ctx replay.RunContext
}

func (f helloFlowImpl) Sleep(d time.Duration) {
	f.ctx.Sleep(d)
}

func (f helloFlowImpl) CreateEvent() *reflex.Event {
	return f.ctx.CreateEvent()
}

func (f helloFlowImpl) LastEvent() *reflex.Event {
	return f.ctx.LastEvent()
}

func (f helloFlowImpl) Run() string {
	return f.ctx.Run()
}

func (f helloFlowImpl) Restart(message *State) {
	f.ctx.Restart(message)
}

func (f helloFlowImpl) EmitMessage(message *String) {
	f.ctx.EmitOutput(_oHelloMessage, message)
}

// StreamHello returns a stream of replay events for the hello workflow and an optional run.
func StreamHello(cl replay.Client, run string) reflex.StreamFunc {
	return cl.Stream(_ns, _wHello, run)
}

// HandleMessage calls fn and returns true if the event is a message output.
// Use StreamHello to provide the events.
func HandleMessage(e *reflex.Event, fn func(run string, message *String) error) (bool, error) {
	var ok bool
	err := replay.Handle(e,
		replay.HandleSkip(func(namespace, workflow, run string) bool {
			return namespace != _ns || workflow != _wHello
		}),
		replay.HandleOutput(func(namespace, workflow, run string, output string, message proto.Message) error {
			if output != _oHelloMessage {
				return nil
			}
			ok = true
			return fn(run, message.(*String))
		}))
	if err != nil {
		return false, err
	}
	return ok, nil
}

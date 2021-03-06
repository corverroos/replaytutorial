package main

// Code generated by typedreplay. DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/luno/reflex"

	"github.com/corverroos/replay"
	// TODO(corver): Support importing other packages.
)

const (
	_ns         = "11_sharding"
	_wHello     = "hello"
	_aSlowPrint = "slow_print"
)

// RunHello provides a type API for running the hello workflow.
// It returns true on success or false on duplicate calls or an error.
func RunHello(ctx context.Context, cl replay.Client, run string, message *String) (bool, error) {
	return cl.RunWorkflow(ctx, _ns, _wHello, run, message)
}

// RegisterHello registers and starts the hello workflow consumer.
func RegisterHello(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore,
	hello func(helloFlow, *String), opts ...replay.Option) {

	helloFunc := func(ctx replay.RunContext, message *String) {
		hello(helloFlowImpl{ctx}, message)
	}

	copied := append([]replay.Option{replay.WithName(_wHello)}, opts...)

	replay.RegisterWorkflow(getCtx, cl, cstore, _ns, helloFunc, copied...)
}

// RegisterSlowPrint registers and starts the SlowPrint activity consumer.
func RegisterSlowPrint(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore, b Backends, opts ...replay.Option) {

	copied := append([]replay.Option{replay.WithName(_aSlowPrint)}, opts...)

	replay.RegisterActivity(getCtx, cl, cstore, b, _ns, SlowPrint, copied...)
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
	// The event timestamp could be used to reason about run liveliness.
	LastEvent() *reflex.Event

	// Now returns the last event timestamp as the deterministic "current" time.
	// It is assumed the first time this is used in logic it will be very close to correct while
	// producing deterministic logic during bootstrapping.
	Now() time.Time

	// Run returns the run name/identifier.
	Run() string

	// Restart completes the current run iteration and starts a new run iteration with the provided input message.
	// The run state is effectively reset. This is handy to mitigate bootstrap load for long running tasks.
	// It also allows updating the activity logic/ordering.
	Restart(message *String)

	// SlowPrint results in the SlowPrint activity being called asynchronously
	// with the provided parameter and returns the response once available.
	SlowPrint(message *String) *Empty
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

func (f helloFlowImpl) Now() time.Time {
	return f.ctx.LastEvent().Timestamp
}

func (f helloFlowImpl) Run() string {
	return f.ctx.Run()
}

func (f helloFlowImpl) Restart(message *String) {
	f.ctx.Restart(message)
}

func (f helloFlowImpl) SlowPrint(message *String) *Empty {
	return f.ctx.ExecActivity(SlowPrint, message, replay.WithName(_aSlowPrint)).(*Empty)
}

// StreamHello returns a stream of replay events for the hello workflow and an optional run.
func StreamHello(cl replay.Client, run string) reflex.StreamFunc {
	return cl.Stream(_ns, _wHello, run)
}

// HandleHelloRun calls fn if the event is a hello RunCreated event.
// Use StreamHello to provide the events.
func HandleHelloRun(e *reflex.Event, fn func(run string, message *String) error) error {
	return replay.Handle(e,
		replay.HandleSkip(func(namespace, workflow, run string) bool {
			return namespace != _ns || workflow != _wHello
		}),
		replay.HandleRunCreated(func(namespace, workflow, run string, message proto.Message) error {
			return fn(run, message.(*String))
		}),
	)
}

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
	_ns               = "07_makerorder"
	_wMakerOrder      = "maker_order"
	_oMakerOrderTrade = "trade"
	_aPlace           = "place"
	_aMonitor         = "monitor"
	_aCancel          = "cancel"
)

// RunMakerOrder provides a type API for running the maker_order workflow.
// It returns true on success or false on duplicate calls or an error.
func RunMakerOrder(ctx context.Context, cl replay.Client, run string, message *Order) (bool, error) {
	return cl.RunWorkflow(ctx, _ns, _wMakerOrder, run, message)
}

// startReplayLoops registers the workflow and activities for typed workflow functions.
func startReplayLoops(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore, b Backends,
	makerOrder func(makerOrderFlow, *Order)) {

	makerOrderFunc := func(ctx replay.RunContext, message *Order) {
		makerOrder(makerOrderFlowImpl{ctx}, message)
	}
	replay.RegisterWorkflow(getCtx, cl, cstore, _ns, makerOrderFunc, replay.WithName(_wMakerOrder))

	replay.RegisterActivity(getCtx, cl, cstore, b, _ns, Place, replay.WithName(_aPlace))
	replay.RegisterActivity(getCtx, cl, cstore, b, _ns, Monitor, replay.WithName(_aMonitor))
	replay.RegisterActivity(getCtx, cl, cstore, b, _ns, Cancel, replay.WithName(_aCancel))
}

// makerOrderFlow defines a typed API for the maker_order workflow.
type makerOrderFlow interface {

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
	Restart(message *Order)

	// Place results in the Place activity being called asynchronously
	// with the provided parameter and returns the response once available.
	Place(message *Order) *OrderRef

	// Monitor results in the Monitor activity being called asynchronously
	// with the provided parameter and returns the response once available.
	Monitor(message *OrderRef) *OrderState

	// Cancel results in the Cancel activity being called asynchronously
	// with the provided parameter and returns the response once available.
	Cancel(message *OrderRef) *Empty

	// EmitTrade stores the trade output in the event log and returns when successful.
	EmitTrade(message *OrderRef)
}

type makerOrderFlowImpl struct {
	ctx replay.RunContext
}

func (f makerOrderFlowImpl) Sleep(d time.Duration) {
	f.ctx.Sleep(d)
}

func (f makerOrderFlowImpl) CreateEvent() *reflex.Event {
	return f.ctx.CreateEvent()
}

func (f makerOrderFlowImpl) LastEvent() *reflex.Event {
	return f.ctx.LastEvent()
}

func (f makerOrderFlowImpl) Now() time.Time {
	return f.ctx.LastEvent().Timestamp
}

func (f makerOrderFlowImpl) Run() string {
	return f.ctx.Run()
}

func (f makerOrderFlowImpl) Restart(message *Order) {
	f.ctx.Restart(message)
}

func (f makerOrderFlowImpl) Place(message *Order) *OrderRef {
	return f.ctx.ExecActivity(Place, message, replay.WithName(_aPlace)).(*OrderRef)
}

func (f makerOrderFlowImpl) Monitor(message *OrderRef) *OrderState {
	return f.ctx.ExecActivity(Monitor, message, replay.WithName(_aMonitor)).(*OrderState)
}

func (f makerOrderFlowImpl) Cancel(message *OrderRef) *Empty {
	return f.ctx.ExecActivity(Cancel, message, replay.WithName(_aCancel)).(*Empty)
}

func (f makerOrderFlowImpl) EmitTrade(message *OrderRef) {
	f.ctx.EmitOutput(_oMakerOrderTrade, message)
}

// StreamMakerOrder returns a stream of replay events for the maker_order workflow and an optional run.
func StreamMakerOrder(cl replay.Client, run string) reflex.StreamFunc {
	return cl.Stream(_ns, _wMakerOrder, run)
}

// HandleTrade calls fn if the event is a trade output.
// Use StreamMakerOrder to provide the events.
func HandleTrade(e *reflex.Event, fn func(run string, message *OrderRef) error) error {
	return replay.Handle(e,
		replay.HandleSkip(func(namespace, workflow, run string) bool {
			return namespace != _ns || workflow != _wMakerOrder
		}),
		replay.HandleOutput(func(namespace, workflow, run string, output string, message proto.Message) error {
			if output != _oMakerOrderTrade {
				return nil
			}
			return fn(run, message.(*OrderRef))
		}),
	)
}

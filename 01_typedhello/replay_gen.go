package main

import (
	"context"
	"time"

	"github.com/luno/reflex"

	"github.com/corverroos/replay"
)

const (
	_ns     = "01_typedhello"
	_wHello = "hello"
)

func RunHello(ctx context.Context, cl replay.Client, run string, message *String) (bool, error) {
	return cl.RunWorkflow(ctx, _ns, _wHello, run, message)
}

func startReplayLoops(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore, b Backends,
	hello func(helloFlow, *String)) {

	helloFunc := func(ctx replay.RunContext, message *String) {
		hello(helloFlowImpl{ctx}, message)
	}
	replay.RegisterWorkflow(getCtx, cl, cstore, _ns, helloFunc, replay.WithName(_wHello))

}

type helloFlow interface {
	Sleep(d time.Duration)

	CreateEvent() *reflex.Event

	LastEvent() *reflex.Event

	Run() string

	Restart(message *String)
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

func (f helloFlowImpl) Restart(message *String) {
	f.ctx.Restart(message)
}

func StreamHello(cl replay.Client, run string) reflex.StreamFunc {
	return cl.Stream(_ns, _wHello, run)
}

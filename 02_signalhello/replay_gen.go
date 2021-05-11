package main

import (
	"context"
	"time"

	"github.com/corverroos/replay"
	"github.com/golang/protobuf/proto"
	"github.com/luno/reflex"
)

const (
	_ns     = "02_signalhello"
	_wHello = "hello"
)

type helloSignal int

const (
	_sHelloName helloSignal = 1
)

var helloSignalMessages = map[helloSignal]proto.Message{
	_sHelloName: new(String),
}

func (s helloSignal) SignalType() int {
	return int(s)
}

func (s helloSignal) MessageType() proto.Message {
	return helloSignalMessages[s]
}

func SignalHelloName(ctx context.Context, cl replay.Client, run string, message *String, extID string) (bool, error) {
	return cl.SignalRun(ctx, _ns, _wHello, run, _sHelloName, message, extID)
}

func RunHello(ctx context.Context, cl replay.Client, run string, message *Empty) (bool, error) {
	return cl.RunWorkflow(ctx, _ns, _wHello, run, message)
}

func startReplayLoops(getCtx func() context.Context, cl replay.Client, cstore reflex.CursorStore, b Backends,
	hello func(helloFlow, *Empty)) {

	helloFunc := func(ctx replay.RunContext, message *Empty) {
		hello(helloFlowImpl{ctx}, message)
	}
	replay.RegisterWorkflow(getCtx, cl, cstore, _ns, helloFunc, replay.WithName(_wHello))

}

type helloFlow interface {
	Sleep(d time.Duration)

	CreateEvent() *reflex.Event

	LastEvent() *reflex.Event

	Run() string

	Restart(message *Empty)

	AwaitName(d time.Duration) (*String, bool)
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

func (f helloFlowImpl) Restart(message *Empty) {
	f.ctx.Restart(message)
}

func (f helloFlowImpl) AwaitName(d time.Duration) (*String, bool) {
	res, ok := f.ctx.AwaitSignal(_sHelloName, d)
	if !ok {
		return nil, false
	}
	return res.(*String), true
}

func StreamHello(cl replay.Client, run string) reflex.StreamFunc {
	return cl.Stream(_ns, _wHello, run)
}

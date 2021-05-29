// Exercise 08_sparkline guides you through using replay to buffer Luno exchange market
// data and to print a constant window sparkline.
package main

import (
	"context"

	"github.com/corverroos/replay/typedreplay"
	"github.com/luno/fate"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	tut "github.com/corverroos/replaytutorial" // Alias replaytutorial to tut for brevity.
	"github.com/corverroos/replaytutorial/lib/luno"
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

type Backends struct {
	Luno luno.Client
}

// Step 3: Define activity functions

// Fetch fetches the Luno market's last trade.
func Fetch(ctx context.Context, b Backends, f fate.Fate, o *Market) (*Double, error) {
	t, err := b.Luno.Ticker(o.Market)
	if err != nil {
		return nil, err
	}

	return &Double{Value: t.LastTrade}, nil
}

// Step 4: Define the typedreplay namespace

var _ = typedreplay.Namespace{
	Name: "08_sparkline",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "buffer",
			Description: "Buffers a Luno market's last N trades",
			Input:       new(Data),
		},
	},
	Activities: []typedreplay.Activity{
		{
			Name:        "fetch",
			Description: "Fetches the Luno market's last trade",
			Func:        Fetch,
		},
	},
}

// Step 5: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// Step 6: Define the buffer workflow function

func buffer(flow bufferFlow, d *Data) {
	// TODO(you): Implement the buffer workflow function.
	//showme:hidden workflow
}

// Step 7: Define your Main function which is equivalent to a main function, just with some prepared state.

func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, buffer)

	// Start the buy maker order for market XBTZAR
	ok, err := RunBuffer(ctx, s.Replay, "1", &Data{
		Market:    "XBTZAR",
		Size:      10,
		PeriodSec: 1,
	})
	if err != nil {
		return err
	} else if !ok {
		log.Info(ctx, "run already started")
	} else {
		log.Info(ctx, "new run started")
	}

	// TODO(you): Define the RunCreated event consume function that prints the sparkline using the generated HandleBufferRun function.
	consume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		panic("implement me")
		//showme:hidden consume
	}

	// Define and run the reflex spec using the generated StreamMakerOrder function.
	spec := reflex.NewSpec(
		StreamBuffer(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("08_sparkline/spark", consume))

	go rpatterns.RunForever(s.AppCtxFunc, spec)

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 8: Run the program and confirm the same expected output.
//go:generate go run github.com/corverroos/replaytutorial/08_sparkline -db_refresh

// Example output:
// I 14:31:01.766 08_sparkline/main.go:104: new run started
// I 14:31:01.766 08_sparkline/main.go:123: Press Ctrl-C to exit...
// I 14:31:01.768 08_sparkline/main.go:110: [consumer=08_sparkline/spark]
// I 14:31:03.772 08_sparkline/main.go:110: ▄[consumer=08_sparkline/spark]
// I 14:31:05.783 08_sparkline/main.go:110: ▁█[consumer=08_sparkline/spark]
// I 14:31:07.787 08_sparkline/main.go:110: ▁██[consumer=08_sparkline/spark]
// I 14:31:09.795 08_sparkline/main.go:110: ▁███[consumer=08_sparkline/spark]
// I 14:31:11.809 08_sparkline/main.go:110: ▁████[consumer=08_sparkline/spark]
// I 14:31:13.822 08_sparkline/main.go:110: ▁█████[consumer=08_sparkline/spark]
// I 14:31:15.838 08_sparkline/main.go:110: ▁██████[consumer=08_sparkline/spark]
// I 14:31:17.860 08_sparkline/main.go:110: ▁███████[consumer=08_sparkline/spark]
// I 14:31:19.880 08_sparkline/main.go:110: ▁████████[consumer=08_sparkline/spark]
// I 14:31:21.906 08_sparkline/main.go:110: ▁█████████[consumer=08_sparkline/spark]
// I 14:31:23.927 08_sparkline/main.go:110: ███████▁▅▅[consumer=08_sparkline/spark]
// I 14:31:25.944 08_sparkline/main.go:110: ██████▁▅▅█[consumer=08_sparkline/spark]
// I 14:31:27.966 08_sparkline/main.go:110: █████▁▅▅██[consumer=08_sparkline/spark]
// I 14:31:29.045 g/c/replaytutorial/setup.go:64: app context canceled

// Step 9: Experiments
// - Does it match up with recent trades on https://www.luno.com/trade/markets/XBTZAR
// - How fast can it go?
// - What happens if you change the period and size?
// - How would you stop the workflow (timeout/deadline/max periods)?

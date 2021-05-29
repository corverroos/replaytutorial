// Exercise 07_makerorder guides you through using replay to place and update a marker order on the Lino exchange.
//
// A maker order is either a limit buy order below the market price or a limit sell order above the market price.
// The order waits in the order book and is therefore said to ‘make’ the market. The Luno exchange,
// https://www.luno.com/trade/markets, charges significantly less maker fees than taker fees.
//
// Maker orders therefore do not trade immediately but waits until the market moves over the maker order price resulting in a trade.
// If the market however moves in the other direction away from the order, the maker order can be cancelled and re-created closer to the
// current price.
package main

import (
	"context"
	"time"

	"github.com/corverroos/replay/typedreplay"
	"github.com/google/uuid"
	"github.com/luno/fate"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	tut "github.com/corverroos/replaytutorial" // Alias replaytutorial to tut for brevity.
	"github.com/corverroos/replaytutorial/lib/luno"
)

// Increase showme to 1 to show next 1 hidden tip
//go:generate go run ../lib/showme -hide 0

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

// Place places a bid or ask post only order at the current market rate and returns the order external id.
func Place(ctx context.Context, b Backends, f fate.Fate, o *Order) (*OrderRef, error) {
	t, err := b.Luno.Ticker(o.Market)
	if err != nil {
		return nil, err
	}

	price := t.Ask
	if o.IsBuy {
		price = t.Bid
	}

	extID := uuid.New().String()

	err = b.Luno.PlacePostOnly(o.Market, o.IsBuy, o.Amount, price, extID)
	if err != nil {
		return nil, err
	}

	log.Info(ctx, "post only order placed", j.MKV{"price": price})

	return &OrderRef{
		Market:        o.Market,
		IsBuy:         o.IsBuy,
		Price:         price,
		ExtId:         extID,
		TimeUnixMilli: time.Now().Unix() * 1000,
	}, nil
}

// Monitor return the order status and the current market price.
func Monitor(ctx context.Context, b Backends, f fate.Fate, o *OrderRef) (*OrderState, error) {
	// TODO(you): Implement the monitor activity, also include some logging
	//showme:hidden monitor
	panic("implement me")
}

// Cancel cancels the limit order.
func Cancel(ctx context.Context, b Backends, f fate.Fate, o *OrderRef) (*Empty, error) {
	log.Info(ctx, "cancelling order")
	return new(Empty), b.Luno.Cancel(o.ExtId)
}

// Step 4: Define the typedreplay namespace

var _ = typedreplay.Namespace{
	Name: "07_makerorder",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "maker_order",
			Description: "Places and maintains a (post-only) maker order on the Luno exchange",
			Input:       new(Order),
			Outputs: []typedreplay.Output{
				{
					Name:        "trade",
					Description: "Reference to traded order",
					Message:     new(OrderRef),
				},
			},
		},
	},
	Activities: []typedreplay.Activity{
		{
			Name:        "place",
			Description: "Places a post-only limit order on the Luno exchange",
			Func:        Place,
		},
		{
			Name:        "monitor",
			Description: "Monitor returns the order state and current market price",
			Func:        Monitor,
		},
		{
			Name:        "cancel",
			Description: "Cancels the placed order",
			Func:        Cancel,
		},
	},
}

// Step 5: Generate the type-safe replay API for the above definition.
//go:generate typedreplay

// Step 6: Define the makerOrder workflow function

// makerOrder workflow function:
// - place the order
// - monitor every 1 sec
// - if Traded, emit output and return
// - else if market moved more than a factor of 0.00001, cancel and restart
// - else monitor again
func makerOrder(flow makerOrderFlow, o *Order) {
	// TODO(you): Implement the makerOrder workflow function.
	//showme:hidden workflow
}

// Step 7: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	// Call the generated startReplayLoops.
	startReplayLoops(s.AppCtxFunc, s.Replay, s.Cursors, Backends{}, makerOrder)

	// Start the buy maker order for market XBTZAR
	ok, err := RunMakerOrder(ctx, s.Replay, "1", &Order{
		Market: "XBTZAR",
		Amount: 10,
		IsBuy:  true,
	})
	if err != nil {
		return err
	} else if !ok {
		log.Info(ctx, "run already started")
	} else {
		log.Info(ctx, "new run started")
	}

	// Define the output consume function
	consume := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		return HandleTrade(e, func(run string, msg *OrderRef) error {
			log.Info(ctx, "maker order traded", j.MKV{"price": msg.Price})
			return nil
		})
	}

	// Define and run the reflex spec using the generated StreamMakerOrder function.
	spec := reflex.NewSpec(
		StreamMakerOrder(s.Replay, ""),
		s.Cursors,
		reflex.NewConsumer("07_makertrade/trades", consume))

	go rpatterns.RunForever(s.AppCtxFunc, spec)

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 8: Run the program and confirm the same expected output.
//go:generate go run github.com/corverroos/replaytutorial/07_makertrade -db_refresh

// Example output:
// I 13:00:51.601 07_makertrade/main.go:198: new run started
// I 13:00:51.602 07_makertrade/main.go:217: Press Ctrl-C to exit...
// I 13:00:52.384 07_makertrade/main.go:65: post only order placed[consumer=replay_activity/07_makerorder/place,price=507063,replay_run=1]
// I 13:00:54.848 07_makertrade/main.go:96: order resting[consumer=replay_activity/07_makerorder/monitor,current_price=507070,diff=7,factor=-1.3804800126249184e-05,replay_run=1]
// I 13:00:54.856 07_makertrade/main.go:108: cancelling order[consumer=replay_activity/07_makerorder/cancel,replay_run=1]
// I 13:00:55.484 07_makertrade/main.go:65: post only order placed[consumer=replay_activity/07_makerorder/place,price=507072,replay_run=1]
// I 13:00:57.408 07_makertrade/main.go:96: order resting[consumer=replay_activity/07_makerorder/monitor,current_price=507075,diff=3,factor=-5.916284573248554e-06,replay_run=1]
// I 13:00:59.411 07_makertrade/main.go:96: order resting[consumer=replay_activity/07_makerorder/monitor,current_price=507081,diff=9,factor=-1.77486437078489e-05,replay_run=1]
// .....
// I 13:01:32.475 07_makertrade/main.go:65: post only order placed[consumer=replay_activity/07_makerorder/place,price=507267,replay_run=1]
// I 13:01:34.327 07_makertrade/main.go:94: order traded[consumer=replay_activity/07_makerorder/monitor,current_price=507268,replay_run=1]
// I 13:01:34.333 07_makertrade/main.go:204: maker order traded[consumer=07_makertrade/trades,price=507267]

// Step 9: Experiments
// - Does it match up with live order book and recent trades on https://www.luno.com/trade/markets/XBTZAR
// - Do a sell maker order
// - What happens if you increase/decrease the restart factor?

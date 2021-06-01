// Exercise 11_sharding guides you through using sharded activity consumers
// the parallelize execution of activity functions. Note that same principal
// applies to workflow consumers.
//
// Activity sharding is useful in multiple use-cases:
// - Parallelize activity functions for improved throughput.
// - Collocating specific runs and specific resources.
// - Dedicated execution of specific long-lived runs.
//
// Workflow sharding is not that useful as workflow functions do not
// perform IO (outside of replay itself).
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/corverroos/replay/typedreplay"
	"github.com/luno/fate"
	"github.com/luno/jettison/log"

	tut "github.com/corverroos/replaytutorial"
)

// Increase showme to 1 to show next hidden solution.
//go:generate go run ../lib/showme 0

// Step 0: main functions always just call tut.Main(Main).
func main() {
	tut.Main(Main)
}

// Step 1: Replay always requires protobufs, so generate your types.
//go:generate protoc --go_out=plugins=grpc:. ./pb.proto

// Step 2: typedreplay requires a locally defined Backends dependency wrapper type.

type Backends struct {
	tut.State
}

// Step 3: Define activity and typedreplay namespace

// SlowPrint prints the provided message.
func SlowPrint(ctx context.Context, b Backends, f fate.Fate, msg *String) (*Empty, error) {
	time.Sleep(time.Second * 2)
	log.Info(ctx, msg.Value)
	return new(Empty), nil
}

var _ = typedreplay.Namespace{
	Name: "11_sharding",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow just prints 'Hello {name}'",
			Input:       new(String),
		},
	},
	Activities: []typedreplay.Activity{
		{
			Name:        "slow_print",
			Description: "Prints the provided message slowly",
			Func:        SlowPrint,
		},
	},
	// ExposeRegisterFuncs so we get access to individual Register functions instead of single `startReplayLoops`.
	ExposeRegisterFuncs: true,
}

// Step 4: Generate the typedreplay API and define the workflow function.

//go:generate typedreplay

func Hello(flow helloFlow, name *String) {
	msg := fmt.Sprintf("Hello %s", name.Value)
	flow.SlowPrint(&String{Value: msg})
}

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {

	// TODO(you): Create 5 new runs and register 5 activity shards, each run should be executed by a different
	//  shard using the WithShard option.
	//showme:hidden run

	log.Info(ctx, "Press Ctrl-C to exit...")
	<-ctx.Done()
	return nil
}

// Step 6: Executed the programs and confirm the same expected output
//go:generate go run github.com/corverroos/replaytutorial/11_sharding

// Expected output of "c":
// I 17:12:53.779 11_sharding/main.go:106: Press Ctrl-C to exit...
// I 17:12:55.787 11_sharding/main.go:50: Hello shard_1[consumer=replay_activity/11_sharding/slow_print/shard_1,replay_run=1622560373|1]
// I 17:12:55.787 11_sharding/main.go:50: Hello shard_0[consumer=replay_activity/11_sharding/slow_print/shard_0,replay_run=1622560373|0]
// I 17:12:55.787 11_sharding/main.go:50: Hello shard_2[consumer=replay_activity/11_sharding/slow_print/shard_2,replay_run=1622560373|2]
// I 17:12:55.790 11_sharding/main.go:50: Hello shard_3[consumer=replay_activity/11_sharding/slow_print/shard_3,replay_run=1622560373|3]
// I 17:12:55.798 11_sharding/main.go:50: Hello shard_4[consumer=replay_activity/11_sharding/slow_print/shard_4,replay_run=1622560373|4]

// Experiments:
// - What happens if you use WithHashShard instead?
// - Try sharding the workflow function as well

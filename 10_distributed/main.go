// Exercise 10_distributed guides you through using the replay in a distributed system where
// creating runs, workflow consumers and activity consumers and event consumers are being
// processed/called from different processes. Each process directly communicates with the
// replay client (in this case the replay DBClient) treating it almost like a common event bus.
//
// To participate, each process just needs access to the typedreplay generated API which defines the namespace
// names and types as well as access to the same replay backend.
//
// The workflow is the same "hello world" example as 00/01. The four processes are:
// - r: Creates new workflow runs.
// - a: Registers and processes the activity function.
// - w: Registers and processes the workflow function.
// - c: Consumes and logs events.
//
// To start a process, provide it's name as argument to this main program.
package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/corverroos/replay/typedreplay"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"

	tut "github.com/corverroos/replaytutorial"
)

// Increase showme to 1 to show next hidden solution.
//go:generate go run ../lib/showme -hide 0

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

// Print prints the provided message.
func Print(ctx context.Context, b Backends, f fate.Fate, msg *String) (*Empty, error) {
	log.Info(ctx, msg.Value)
	return new(Empty), nil
}

var _ = typedreplay.Namespace{
	Name: "10_distributed",
	Workflows: []typedreplay.Workflow{
		{
			Name:        "hello",
			Description: "Hello workflow just prints 'Hello {name}'",
			Input:       new(String),
		},
	},
	Activities: []typedreplay.Activity{
		{
			Name:        "print",
			Description: "Prints the provided message",
			Func:        Print,
		},
	},
	// ExposeRegisterFuncs so we get access to individual Register functions instead of single `startReplayLoops`.
	ExposeRegisterFuncs: true,
}

// Step 4: Generate the typedreplay API and define the workflow function.

//go:generate typedreplay

func Hello(flow helloFlow, name *String) {
	flow.Sleep(time.Second * 2)

	msg := fmt.Sprintf("Hello %s", name.Value)

	flow.Print(&String{Value: msg})
}

// ttl defines how long a process should live before exiting.
var ttl = flag.Duration("ttl", time.Minute, "TTL of the process")

// Step 5: Define your Main function which is equivalent to a main function, just with some prepared state.
func Main(ctx context.Context, s tut.State) error {
	if len(flag.Args()) != 1 {
		return errors.New("please provide the process name as single argument: go run thispkg a/r/w/c")
	}

	proc := flag.Arg(0)

	switch proc {
	case "r":
		// TODO(you): The r process should just create a new workflow run and then exit.
		//showme:hidden r
	case "a":
		// TODO(you): The a process should register the Print activity and then block below until TTL is reached.
		//showme:hidden a
	case "w":
		// TODO(you): The w process should register the Hello workflow and then block below until TTL is reached.
		//showme:hidden w
	case "c":
		// TODO(you): The c process should consume and consume and RunCreated and RunCompleted events.
		//showme:hidden c
	}

	log.Info(ctx, "process started", j.MKV{"proc": proc, "ttl": *ttl})
	select {
	case <-time.After(*ttl):
		log.Info(ctx, "process ttl reached")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Step 6: Executed the programs and confirm the same expected output
//go:generate go run github.com/corverroos/replaytutorial/10_distributed -server_loops=true c
//go:generate go run github.com/corverroos/replaytutorial/10_distributed -server_loops=false w
//go:generate go run github.com/corverroos/replaytutorial/10_distributed -server_loops=false a

// Executed this multiple times to create multiple runs.
//go:generate go run github.com/corverroos/replaytutorial/10_distributed -server_loops=false r

// Expected output of "c":
// I 17:41:46.005 10_distributed/replay.go:138: process started[proc=c,ttl=1m0s]
// I 17:41:58.583 10_distributed/replay.go:121: run created[consumer=10_distributed,run=09f9b171-b47d-4cce-9c82-da1a2b493802]
// I 17:42:01.073 10_distributed/replay.go:125: run completed[consumer=10_distributed,run=09f9b171-b47d-4cce-9c82-da1a2b493802]
// I 17:42:07.227 10_distributed/replay.go:121: run created[consumer=10_distributed,run=97506cfc-951f-48c8-afc4-18976f451e03]
// I 17:42:10.109 10_distributed/replay.go:125: run completed[consumer=10_distributed,run=97506cfc-951f-48c8-afc4-18976f451e03]

// Expected output of "a":
// I 17:41:55.236 10_distributed/replay.go:138: process started[proc=a,ttl=1m0s]
// I 17:42:01.069 10_distributed/replay.go:56: Hello world[consumer=replay_activity/10_distributed/print,replay_run=09f9b171-b47d-4cce-9c82-da1a2b493802]
// I 17:42:10.102 10_distributed/replay.go:56: Hello world[consumer=replay_activity/10_distributed/print,replay_run=97506cfc-951f-48c8-afc4-18976f451e03]

// Experiments:
// - What happens if either w or a is not running?
// - Why should only one process run the server loops?

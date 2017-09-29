package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	notary "github.com/NTRYPlatform/ntry-backend"

	docopt "github.com/docopt/docopt-go"
)

/**
 * Expected flow
 * Connect to datebase when program start (expensive)
 * Disconnect when the program is terminated
 * Create a session per request (cheap)
 * Clean up after each request
 */
/** HTTP Middleware
 * - Pluggable and self-contained piece of code that wraps web application handlers
 * - Componentst that work as another layer in the request handling cycle, which can
 * execute some logic before or after executing your http application handlers
 * - Great for implementing cross-cuttin concerns: Uthentication, authorization, caching,
 * logging etc.
 */

var usage = `
Usage:
    server -c <confpath> [options] 
    server -h | --help
Options:
    -l --logpath=<path>     Log file path [default: /dev/null].
    -c --confpath=<path>    Configuration file path [default: .notaryconf/ntryapp.yml].
    -d --debug              Enable debug mode [default: false].
    -h --help               Show this screen.
`

func onErrorExit(err error) {
	if err != nil {
		fmt.Printf("[server  ] %+v\n", err)
		os.Exit(0)
	}
}

func onInterruptSignal(fn func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		fn()
	}()
}

func main() {

	args, err := docopt.Parse(usage, os.Args[1:], true, "", true)
	onErrorExit(err)

	ntry, err := notary.New(args)
	onErrorExit(err)

	err = ntry.Init()
	onErrorExit(err)

	exitMtx := sync.RWMutex{}
	exit := func() {
		// See if there was a panic...
		fmt.Fprintln(os.Stderr, recover())
		exitMtx.Lock()
		println() // make it look nice after the ^C
		fmt.Println("[notary  ] shutting down...")
		onErrorExit(ntry.Shutdown())
	}
	defer exit()

	onInterruptSignal(func() {
		exit()
		os.Exit(0)
	})

	onErrorExit(ntry.Start())

}

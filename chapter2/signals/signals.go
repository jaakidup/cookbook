package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig sets up a listener for
// SIGINT interrupts
func CatchSig(ch chan os.Signal, done chan bool) {
	// block on waiting for signal
	sig := <-ch
	// print it when it's received
	fmt.Println("signal recieved: ", sig)

	// we can set up handlers for all types of signals

	switch sig {
	case syscall.SIGINT:
		fmt.Println("Handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("Handling a SIGTERM now!")
	default:
		fmt.Println("Unexpected signal received")
	}
	// terminate
	done <- true
}

func main() {
	// initialize our channels
	signals := make(chan os.Signal)
	done := make(chan bool)

	// hook them up to the signals lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// if a signal is caugth by this go routine
	// it wil write to done
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}

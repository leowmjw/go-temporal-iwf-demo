package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Welcome to Demo Subscription ..")
	Run()
}

func Run() {
	fmt.Println("start running samples")
	closeFn := startIWFWorker()
	// Clean up iWF
	defer closeFn()
	//// We can run normal Temporal workers too ..
	//closeTemporalFn := startTemporalWorker()
	//defer closeTemporalFn()

	// Block till SIGTERM ..
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	s := <-c
	fmt.Println("GOT_SIG:", s.String())
	fmt.Println("UNBLOCKED ..")
	// At this point the defers should kick in ..
}

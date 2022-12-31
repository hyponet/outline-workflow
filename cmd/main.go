package main

import (
	"context"
	"github.com/hyponet/outline-workflow/cmd/apps"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	stopCh := SetupSignalHandler()
	ctx, canF := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		canF()
	}()

	if err := apps.RootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal

func SetupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

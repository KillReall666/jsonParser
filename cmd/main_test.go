package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignalHandling(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	go func() {
		<-signals
		cancel()
	}()

	go func() {
		signals <- syscall.SIGTERM
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ctx.Done():
		t.Log("Context was cancelled as expected")
	default:
		t.Fatal("Context was not cancelled as expected")
	}
}

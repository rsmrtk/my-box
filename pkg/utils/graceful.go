package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// GracefulShutdown listens for an interrupt signal and returns a channel that will be closed when the signal is received.
func GracefulShutdown() <-chan struct{} {
	shutdown := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)

		<-s
		close(shutdown)
	}()
	return shutdown
}

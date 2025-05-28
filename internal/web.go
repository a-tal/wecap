package internal

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ll          = log.New(os.Stdout, "[wecap] ", log.Default().Flags())
	handlerPath = "/data/report/"
)

// Run the wecap web server until we notice a signal
func Run() {
	srv := newServer()

	// run until sig int|term
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	// give shutdown 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		ll.Printf("failed to shutdown cleanly: %v", err)
	}
}

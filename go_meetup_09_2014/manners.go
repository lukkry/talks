package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
)

func main() {
	handler := MyHTTPHandler()

	server := manners.NewServer()
	go catchTermSignal(server)

	server.ListenAndServe(":7000", handler)
}

func catchTermSignal(server *manners.GracefulServer) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch
	server.Shutdown <- true
}

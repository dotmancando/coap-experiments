package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	server := &http.Server{
		Addr:    ":6000",
		Handler: mux,
	}

	go func(s *http.Server) {
		sig := <-gracefulStop
		fmt.Printf("caught signal: %+v\n", sig)
		s.Shutdown(nil)
		os.Exit(0)
	}(server)

	server.ListenAndServe()
}

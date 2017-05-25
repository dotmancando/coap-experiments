package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func pulse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	w.Write(body)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			log.Printf("[%s] %s. Completed in %v", r.Method, r.URL.String(), time.Now().Sub(startTime))
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	mux := http.NewServeMux()
	mux.Handle("/pulse", loggingMiddleware(http.HandlerFunc(pulse)))
	mux.Handle("/echo", loggingMiddleware(http.HandlerFunc(echo)))

	addr := ":6000"

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func(s *http.Server) {
		sig := <-gracefulStop
		fmt.Printf("Caught signal: %+v\n", sig)
		s.Shutdown(nil)
		os.Exit(0)
	}(server)

	log.Printf("Started HTTP server listening on %s", addr)
	server.ListenAndServe()
}

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

var app = NewApplication()

var router = httprouter.New()

var client = &http.Client{}

func main() {
	log.Println("Starting web service")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	flag.UintVar(&app.Capacity, "capacity", 10000000, "Capacity for the Cuckoo Filter in bytes")
	flag.StringVar(&app.Storage, "storage", "storage.db", "Filename with a copy of the Cuckoo Filter")
	flag.StringVar(&app.Hostname, "hostname", "babywaf.test", "Domain where the web service is running")
	flag.StringVar(&app.AuthSecret, "authSecret", "LRvHBZG5m4Xfj9RuWQJ0cVbRBg7uRENBm7UzLD6X", "Shared secret key to communicate with other servers")
	flag.StringVar(&app.ListenAddr, "listenAddr", ":80", "Hostname and port where the server is listening")
	flag.DurationVar(&app.ReadTimeout, "readTimeout", time.Second*5, "Maximum amount of time to read the request body")
	flag.DurationVar(&app.ReadHeaderTimeout, "readHeaderTimeout", time.Second*5, "Maximum amount of time to read the request headers")
	flag.DurationVar(&app.WriteTimeout, "writeTimeout", time.Second*5, "Maximum amount of time to write the response")
	flag.DurationVar(&app.IdleTimeout, "idleTimeout", time.Second*5, "Maximum amount of time to wait for the client")
	flag.DurationVar(&app.RequestTimeout, "requestTimeout", time.Second*5, "Maximum amount of time to wait for internal requests")
	flag.DurationVar(&app.ShutdownTimeout, "shutdownTimeout", time.Second*5, "Maximum amount of time to wait for the server shutdown")
	flag.IntVar(&app.ServerPrefixN, "serverPrefixN", 1, "Number of characters from the URL-hash to build server pool")

	flag.Parse()

	app.LoadDatabase()

	server := &http.Server{
		Addr:              app.ListenAddr,
		Handler:           router,
		ReadTimeout:       app.ReadTimeout,
		ReadHeaderTimeout: app.ReadHeaderTimeout,
		WriteTimeout:      app.WriteTimeout,
		IdleTimeout:       app.IdleTimeout,
	}

	// NOTES(yorman): try to keep the request timeout lower or equal to the
	// server write timeout. This allows the web service to finish operations
	// in the queue that are taking longer than expected.
	client.Timeout = app.RequestTimeout

	go func() {
		log.Println("Listening on " + app.ListenAddr)

		if err := server.ListenAndServe(); err != nil {
			log.Println("http.ListenAndServe", err)
		}
	}()

	// NOTES(yorman): implement graceful server shutdown. Make sure to export
	// the data collected inside the Cuckoo Filter to be able to recover the
	// unsafe URLs during subsequent server restarts.
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), app.ShutdownTimeout)
	defer cancel()
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Println("server.Shutdown", err)
	}

	log.Println("finished")
}

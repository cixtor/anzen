package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var app = NewApplication()

var router = httprouter.New()

func main() {
	fmt.Println("Starting web service")

	flag.UintVar(&app.Capacity, "capacity", 10000000, "Capacity for the Cuckoo Filter in bytes")
	flag.StringVar(&app.Storage, "storage", "storage.db", "Filename with a copy of the Cuckoo Filter")
	flag.StringVar(&app.ListenAddr, "listenAddr", ":8080", "Hostname and port where the server is listening")
	flag.DurationVar(&app.ReadTimeout, "readTimeout", time.Second*5, "Maximum amount of time to read the request body")
	flag.DurationVar(&app.ReadHeaderTimeout, "readHeaderTimeout", time.Second*5, "Maximum amount of time to read the request headers")
	flag.DurationVar(&app.WriteTimeout, "writeTimeout", time.Second*5, "Maximum amount of time to write the response")
	flag.DurationVar(&app.IdleTimeout, "idleTimeout", time.Second*5, "Maximum amount of time to wait for the client")

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

	fmt.Println("Listening on " + app.ListenAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Println("http.ListenAndServe", err)
	}
}

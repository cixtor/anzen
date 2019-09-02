package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var app = NewApplication()

var router = httprouter.New()

func main() {
	fmt.Println("Starting web service")

	flag.UintVar(&app.Capacity, "capacity", 10000000, "Capacity for the Cuckoo Filter in bytes")
	flag.StringVar(&app.Storage, "storage", "storage.db", "Filename with a copy of the Cuckoo Filter")

	flag.Parse()

	app.LoadDatabase()

	fmt.Println("Listening on http://0.0.0.0:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("http.ListenAndServe", err)
	}
}

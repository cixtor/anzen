package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

func main() {
	fmt.Println("Listening on http://0.0.0.0:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("http.ListenAndServe", err)
	}
}

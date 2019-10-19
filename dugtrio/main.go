package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const secret string = "LRvHBZG5m4Xfj9RuWQJ0cVbRBg7uRENBm7UzLD6X"

var router = httprouter.New()

var fakeRedis = NewFakeRedis()

func main() {
	router.POST("/api/insert/:hash", fakeRedis.Insert)
	router.GET("/api/retrieve/:hash", fakeRedis.Retrieve)

	log.Println("Threat-Info server is ready")

	if err := http.ListenAndServe(":80", router); err != nil {
		log.Println("http.ListenAndServe", err)
	}
}

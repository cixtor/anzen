package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/urlinfo/1/:host/:path", urlinfo)
}

func urlinfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	host := ps.ByName("host")
	path := ps.ByName("path")

	if _, err := w.Write(ThreatLevel(host, path)); err != nil {
		log.Println("urlinfo", "w.Write", err)
	}
}

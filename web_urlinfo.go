package main

import (
	"encoding/json"
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

	info, err := ThreatType(host, path)

	if err != nil {
		log.Println("urlinfo", "ThreatType", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Println("urlinfo", "json.Encode", err)
	}
}

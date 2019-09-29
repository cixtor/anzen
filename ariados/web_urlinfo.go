package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/urlinfo/1/*query", urlinfo)
}

func urlinfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var query string
	var info ThreatInfo

	if query, err = TargetURL(ps); err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if info, err = app.ThreatType(query); err != nil {
		log.Println("urlinfo", "ThreatType", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(info); err != nil {
		log.Println("urlinfo", "json.Encode", err)
		return
	}
}

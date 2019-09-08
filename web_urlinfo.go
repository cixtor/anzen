package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/urlinfo/1/*query", urlinfo)
}

func urlinfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var query string
	var info ThreatInfo

	if query = ps.ByName("query"); query == "" || query == "/" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// NOTES(yorman): when the HTTP router selects the named parameter "query"
	// it returns the character directly adjacent to the static URL, in this
	// case a forward slash after "/urlinfo/1"; I expected the router to return
	// the characters after the forward slash considering there is an implicit
	// redirection from "/urlinfo/1" to "/urlinfo/1/" ¯\_(ツ)_/¯
	query = strings.TrimLeft(query, "/")

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

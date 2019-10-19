package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/insert/1/*query", insert)
}

func insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var query string

	if query, err = TargetURL(ps); err != nil {
		log.Println("insert", "TargetURL", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if app.Database.Insert(HashURL(query)) {
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Println("insert", "Database", "failure", query)
	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
}

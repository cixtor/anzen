package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/", index)
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "https://bitbucket.org/cixtor/babywaf", http.StatusFound)
}

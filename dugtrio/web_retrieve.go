package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (redis *FakeRedis) Retrieve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var hash string

	// NOTES(yorman): rudimentary API authentication system.
	if key := r.Header.Get("X-Auth-Secret"); key != secret {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
		return
	}

	if hash = ps.ByName("hash"); hash == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// NOTES(yorman): find the hash in the malware database, if possible.
	info, exists := redis.Entries[hash]

	if !exists {
		info = ThreatInfo{Threat: "NONE"}
	}

	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Println("FakeRedis", "Insert", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
}

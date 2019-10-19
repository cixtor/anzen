package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (redis *FakeRedis) Retrieve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var hash string
	var exists bool
	var info ThreatInfo

	// NOTES(yorman): rudimentary API authentication system.
	if key := r.Header.Get("X-Auth-Secret"); key != secret {
		log.Println("retrieve", "x-auth-secret", "invalid")
		http.Error(w, http.StatusText(403), http.StatusForbidden)
		return
	}

	if hash = ps.ByName("hash"); hash == "" {
		log.Println("retrieve", "url-hash", "missing")
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// NOTES(yorman): find the hash in the malware database, if possible.
	if info, exists = redis.Entries[hash]; exists {
		log.Println("retrieve", "Database", "found")
	} else {
		log.Println("retrieve", "Database", "notfound")
		info = ThreatInfo{Threat: "NONE"}
	}

	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Println("retrieve", "Database", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
}

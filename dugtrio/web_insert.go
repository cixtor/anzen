package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (redis *FakeRedis) Insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var hash string
	var info ThreatInfo

	// NOTES(yorman): rudimentary API authentication system.
	if key := r.Header.Get("X-Auth-Secret"); key != secret {
		log.Println("insert", "x-auth-secret", "invalid")
		http.Error(w, http.StatusText(403), http.StatusForbidden)
		return
	}

	if hash = ps.ByName("hash"); hash == "" {
		log.Println("insert", "url-hash", "missing")
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// NOTES(yorman): expecting JSON({threat:string, platform:string, url:string})
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		log.Println("insert", "Database", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// NOTES(yorman): for the sake of simplicity I am omitting important checks
	// to validate the integrity of the data. For example, is the threat type
	// valid? If not, then choose "UNSPECIFIED", same thing for the platform.
	// For now, trust that whoever has the API key is a legitimate client.

	// NOTES(yorman): insert only if the URL hash is unique.
	if _, ok := redis.Entries[hash]; !ok {
		redis.Lock()
		redis.Entries[hash] = info
		redis.Unlock()
	}

	log.Println("insert", "Database", "success")

	w.WriteHeader(http.StatusOK)
}

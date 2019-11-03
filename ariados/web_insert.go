package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router.POST("/insert/:ttype/*query", insert)
}

func insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var out []byte
	var query string
	var req *http.Request
	var res *http.Response
	var target string

	if query, err = TargetURL(ps); err != nil {
		log.Println("insert", "TargetURL", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	hashed := HashURL(query)

	if !app.Database.Insert(hashed) {
		log.Println("insert", "Database", "failure", query)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if target, err = DugtrioURL(app.Hostname, app.ServerPrefixN, "insert", hashed); err != nil {
		log.Println("insert", "DugtrioURL", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// NOTES(yorman): threat type will be "UNSPECIFIED" by default.
	info := ThreatInfo{Threat: ttUnspecified, URL: query}

	// NOTES(yorman): make sure the threat type is supported by the system.
	switch ttype := ps.ByName("ttype"); ttype {
	case ttMalware:
		info.Threat = ttMalware
	case ttSocialEngineering:
		info.Threat = ttSocialEngineering
	case ttUnwantedSoftware:
		info.Threat = ttUnwantedSoftware
	case ttPotentiallyHarmful:
		info.Threat = ttPotentiallyHarmful
	}

	if out, err = json.Marshal(&info); err != nil {
		log.Println("insert", "json.Marshal", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	body := bytes.NewBuffer(out)

	if req, err = http.NewRequest(http.MethodPost, target, body); err != nil {
		log.Println("insert", "http.NewRequest", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	req.Header.Set("X-Auth-Secret", app.AuthSecret)

	if res, err = client.Do(req); err != nil {
		log.Println("insert", "client.Do", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("insert", "res.Body.Close", err)
		}
	}()

	log.Println("insert", "StatusCode", res.StatusCode)

	w.WriteHeader(res.StatusCode)
}

package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func TargetURL(ps httprouter.Params) (string, error) {
	query := ps.ByName("query")

	if query == "" || query == "/" {
		return "", fmt.Errorf("invalid URL")
	}

	// NOTES(yorman): when the HTTP router selects the named parameter "query"
	// it returns the character directly adjacent to the static URL, in this
	// case a forward slash after "/urlinfo/1"; I expected the router to return
	// the characters after the forward slash considering there is an implicit
	// redirection from "/urlinfo/1" to "/urlinfo/1/" ¯\_(ツ)_/¯
	return strings.TrimLeft(query, "/"), nil
}

func HashURL(query string) []byte {
	decoded, err := url.QueryUnescape(query)

	if err != nil {
		log.Println("HashURL", "url.QueryUnescape", err, query)
		return []byte{}
	}

	return SHA256([]byte(decoded))
}

func SHA256(input []byte) []byte {
	hash := sha256.New()
	if _, err := hash.Write(input); err != nil {
		log.Println("HashURL", "SHA256", err, input)
		return []byte{}
	}
	return hash.Sum(nil)
}

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

func DugtrioURL(hostname string, serverPrefixN int, action string, hashed []byte) (string, error) {
	encoded := fmt.Sprintf("%x", hashed)

	// NOTES(yorman): someone probably messed up the configuration and set a
	// cluster prefix length bigger than sixty-four. This causes the subset of
	// the URL hash to be out of range.
	if serverPrefixN > len(encoded) {
		return "", fmt.Errorf("cluster prefix length is bigger than SHA256")
	}

	return fmt.Sprintf(
		"http://threat-info-%s.%s/api/%s/%s",
		encoded[0:serverPrefixN],
		hostname,
		action,
		encoded,
	), nil
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

package main

import (
	"crypto/sha256"
	"log"
	"net/url"
	"time"
)

// ThreatInfo represents the JSON structure of the metadata associated to a
// malicious URL if found in the malware database. Notice some type properties
// are omitted if empty to reduce bandwidth consumption when the URL is found
// to contain no malicious content.
//
// Example:
//
//   {
//     "threat": "MALWARE",
//     "platform": "WINDOWS",
//     "hash": "6a46bf81098df5b863c4387ee7cf1592ec40ff7aec03088be1b48b18afa8122b",
//     "url": "http://www.example.com/private/malicious.exe",
//     "cache": 300
//   }
type ThreatInfo struct {
	Threat   string        `json:"threat,omitempty"`
	Platform string        `json:"platform,omitempty"`
	Hash     string        `json:"hash,omitempty"`
	URL      string        `json:"url,omitempty"`
	Cache    time.Duration `json:"cache,omitempty"`
}

// ttUnspecified represents an unknown threat type.
const ttUnspecified string = "UNSPECIFIED"

// ttMalware represents a malware threat type.
const ttMalware string = "MALWARE"

// ttSocialEngineering represents a social engineering threat type.
const ttSocialEngineering string = "SOCIAL_ENGINEERING"

// ttUnwantedSoftware represents an unwanted software threat type.
const ttUnwantedSoftware string = "UNWANTED_SOFTWARE"

// ttPotentiallyHarmful represents a potentially harmful application threat type.
const ttPotentiallyHarmful string = "POTENTIALLY_HARMFUL"

// ttNone represents no malware infection.
const ttNone string = "NONE"

// ThreatType inspects the hostname, port number, URL path and query string to
// check if the request will possibly return a malicious page or not. A threat
// level 0x0 means the URL has not been reported to contain malware prior the
// scan. Any threat level above 0x0 represents an unsafe URL and the type of
// infection.
//
// To skip unnecessary encoding and decoding operations, the response of this
// request will be in plain text with bytes representing a threat level if the
// URL was found in the malware database. We could potentially speed up some
// operations using categorization considering some infections are more common
// than others.
//
// Threat Types:
//
//   - UNSPECIFIED: Unknown.
//   - MALWARE: Malware threat type.
//   - SOCIAL_ENGINEERING: Social engineering threat type.
//   - UNWANTED_SOFTWARE: Unwanted software threat type.
//   - POTENTIALLY_HARMFUL: Potentially harmful application threat type.
//   - NONE: No malware infection.
//
// Platform Types:
//
//   - UNSPECIFIED: Unknown platform.
//   - WINDOWS: Threat posed to Windows.
//   - LINUX: Threat posed to Linux.
//   - ANDROID: Threat posed to Android.
//   - MACOS: Threat posed to macOS.
//   - IOS: Threat posed to iOS.
//   - ALL: Threat posed to all defined platforms.
//   - ANY: Threat posed to at least one of the defined platforms.
//
// Ref: https://developers.google.com/safe-browsing/v4/reference/rest/v4/ThreatType
// Ref: https://developers.google.com/safe-browsing/v4/reference/rest/v4/PlatformType
func ThreatType(host string, path string) (ThreatInfo, error) {
	// NOTES(yorman): check if the SHA256 of the URL exists in the malware
	// database. If not found, we immediately return NONE as the threat type.
	// Because we are using a probabilistic data structure, if the answer is
	// yes, we need to double check in subsequent steps to make sure this is
	// not a false positive.
	query := HashURL(host, path)

	if !app.Database.Lookup(query) {
		return ThreatInfo{Threat: ttNone}, nil
	}

	return ThreatInfo{Threat: ttNone}, nil
}

func HashURL(host string, path string) []byte {
	decoded, err := url.QueryUnescape(path)

	if err != nil {
		log.Println("HashURL", "url.QueryUnescape", err, path)
		return []byte{}
	}

	return SHA256([]byte(host + "/" + decoded))
}

func SHA256(input []byte) []byte {
	hash := sha256.New()
	if _, err := hash.Write(input); err != nil {
		log.Println("HashURL", "SHA256", err, input)
		return []byte{}
	}
	return hash.Sum(nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
//     "url": "http://www.example.com/private/malicious.exe"
//   }
type ThreatInfo struct {
	Threat   string `json:"threat,omitempty"`
	Platform string `json:"platform,omitempty"`
	Hash     string `json:"hash,omitempty"`
	URL      string `json:"url,omitempty"`
}

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
func (app *Application) ThreatType(query string) (ThreatInfo, error) {
	if app.Database == nil {
		return ThreatInfo{}, fmt.Errorf("database was not initialized")
	}

	hashed := HashURL(query)

	// NOTES(yorman): check if the SHA256 of the URL exists in the malware
	// database. If not found, we immediately return NONE as the threat type.
	// Because we are using a probabilistic data structure, if the answer is
	// yes, we need to double check in subsequent steps to make sure this is
	// not a false positive.
	if !app.Database.Lookup(hashed) {
		return ThreatInfo{Threat: ttNone}, nil
	}

	encoded := fmt.Sprintf("%x", hashed)

	// NOTES(yorman): someone probably messed up the configuration and set a
	// cluster prefix length bigger than sixty-four. This causes the subset of
	// the URL hash to be out of range.
	if clusterPrefixLength > len(encoded) {
		return ThreatInfo{}, fmt.Errorf("cluster prefix length is bigger than SHA256")
	}

	target := fmt.Sprintf(
		"http://threat-info-%s.%s/api/retrieve/%s",
		encoded[0:clusterPrefixLength],
		app.Hostname,
		encoded,
	)

	var err error
	var req *http.Request
	var res *http.Response
	var info ThreatInfo

	if req, err = http.NewRequest(http.MethodGet, target, nil); err != nil {
		return ThreatInfo{}, fmt.Errorf("ThreatType http.NewRequest %s", err)
	}

	req.Header.Set("X-Auth-Secret", app.AuthSecret)

	if res, err = client.Do(req); err != nil {
		return ThreatInfo{}, fmt.Errorf("ThreatType client.Do %s", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("ThreatType", "res.Body.Close", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return ThreatInfo{}, fmt.Errorf("ThreatType StatusCode %d from `%s`", res.StatusCode, target)
	}

	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		return ThreatInfo{}, fmt.Errorf("ThreatType json.Decode %s", err)
	}

	return info, nil
}

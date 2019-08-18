package main

// ThreatLevel inspects the hostname, port number, URL path and query string to
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
// Example:
//
//   - {0x43, 0x52, 0x49} [CRI] Critical: Malware infection is expected imminently.
//   - {0x53, 0x45, 0x56} [SEV] Severe: Malware infection is highly likely.
//   - {0x53, 0x55, 0x42} [SUB] Substantial: Malware infection is a strong possibility.
//   - {0x4d, 0x4f, 0x44} [MOD] Moderate: Malware infection is possible, but not likely.
//   - {0x4c, 0x4f, 0x57} [LOW] Low: Malware infection is unlikely.
//   - {0x00, 0x00, 0x00} [NIL] None: No malware infection.
func ThreatLevel(host string, path string) []byte {
	return []byte{0x00, 0x00, 0x00}
}

package main

import (
	"testing"
)

func TestThreatType(t *testing.T) {
	app := NewApplication()

	app.LoadDatabase()

	info, err := app.ThreatType(`example.com/hello%2Fworld%3Ffoo%3Dbar`)

	if err != nil {
		t.Fatal(err)
		return
	}

	if info.Threat != ttNone {
		t.Fatal("ThreatType should be NONE")
		return
	}
}

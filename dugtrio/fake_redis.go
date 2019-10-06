package main

import (
	"sync"
)

type FakeRedis struct {
	sync.Mutex
	Entries map[string]ThreatInfo
}

type ThreatInfo struct {
	Threat   string `json:"threat,omitempty"`
	Platform string `json:"platform,omitempty"`
	Hash     string `json:"hash,omitempty"`
	URL      string `json:"url,omitempty"`
}

func NewFakeRedis() *FakeRedis {
	return &FakeRedis{Entries: make(map[string]ThreatInfo)}
}

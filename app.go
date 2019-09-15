package main

import (
	"time"

	cuckoo "github.com/seiflotfy/cuckoofilter"
)

type Application struct {
	// ListenAddr is the hostname and port where the server is listening.
	ListenAddr string

	// Database is where the SHA256 of the malicious URLs are stored.
	Database *cuckoo.Filter

	// Capacity is the size for the Cuckoo Filter in bytes. A capacity of
	// 1,000,000 bytes is a normal default, which allocates about ~1MB on
	// 64-bit machines.
	Capacity uint

	// Storage is the filename with a copy of the Cuckoo Filter. The web
	// service frequently exports the database to this file to maintain the
	// integrity of the database, for example, every time a new malicious URL
	// is added to the blacklist and during the graceful server shutdown. The
	// database is loaded into memory when the server is restarted.
	Storage string

	ReadTimeout time.Duration

	ReadHeaderTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration
}

func NewApplication() *Application {
	return &Application{Database: cuckoo.NewFilter(1000)}
}

func (app *Application) LoadDatabase() {
	app.Database = cuckoo.NewFilter(app.Capacity)
}

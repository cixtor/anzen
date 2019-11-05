package main

import (
	"io/ioutil"
	"log"
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

	// Hostname is the domain associated to the web service. The hostname is
	// also used to construct the URL for the server pool hosting the malware
	// database. For example, if the hostname is example.test then the server
	// pool will look like this: threat-info-PREFIX.example.test
	Hostname string

	// AuthSecret is the shared secret key to communicate with other servers.
	AuthSecret string

	// ServerPrefixN is the number of characters from the URL-hash that will
	// help create the server prefix for the additional web service that we
	// are going to communicate with. The bigger the number
	// the bigger the pool of servers we expect to communicate with.
	//
	// If the length is set to one the size of the server pool will be sixteen,
	// one server for each hexadecimal character [0-9] + [a-f]. If the length
	// is two then the size of the server pool will be 256, one for every
	// combination like so: 00, 01, 02, ... 0a, 0b, 0c and so on until [ff].
	//
	// The maximum size would then be sixty-four which is the size of hashed
	// data using SHA256. Because zero-zero is a valid prefix for a SHA256 sum
	// it means we have to consider repetitions in the calculation of the
	// permutations. So theoretically speaking, the size of the server pool
	// would be equal to:
	//
	//   115,792,089,237,316,...
	//   195,423,570,985,008,...
	//   687,907,853,269,984,...
	//   665,640,564,039,457,...
	//   584,000,000,000,000,000
	//
	// PS: I don’t even know how to read that number :thinking_face:
	ServerPrefixN int

	ReadTimeout time.Duration

	ReadHeaderTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration

	RequestTimeout time.Duration

	ShutdownTimeout time.Duration
}

func NewApplication() *Application {
	return &Application{Database: cuckoo.NewFilter(1000)}
}

func (app *Application) LoadDatabase() {
	var err error
	var out []byte
	var koo *cuckoo.Filter

	// NOTES(yorman): the database was already initialized with 1,000 bytes
	// when the application was created. Override here with the capacity set
	// in the server configuration.
	app.Database = cuckoo.NewFilter(app.Capacity)

	// NOTES(yorman): if the external storage file exists, load its content
	// into the Cuckoo Filter data structure, otherwise print a warning but
	// continue with the rest of our business logic.
	if out, err = ioutil.ReadFile(app.Storage); err != nil {
		log.Println("LoadDatabase", "ioutil.ReadFile", err)
		return
	}

	log.Printf("%d bytes loaded into Cuckoo Filter\n", len(out))

	if koo, err = cuckoo.Decode(out); err != nil {
		log.Println("LoadDatabase", "cuckoo.Decode", err)
		return
	}

	app.Database = koo
}

func (app *Application) ExportDatabase() {
	out := app.Database.Encode()

	if err := ioutil.WriteFile(app.Storage, out, 0644); err != nil {
		log.Println("ExportDatabase", "ioutil.WriteFile", err)
		return
	}

	log.Println("ExportDatabase", "ioutil.WriteFile", "success")
}

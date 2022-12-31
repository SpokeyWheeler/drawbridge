package main

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/knadh/koanf"
	//"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

var db *sql.DB

func OpenDb(dbname string) *sql.DB {
	ConnectionURL := fmt.Sprintf("postgres://spokey@localhost:5432/%s", dbname)
	
}

func main() {
	f := file.Provider("drawbridge.yaml")
	if err := k.Load(f, yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	k.Print()
	log.Println("permission's name is = ", k.String("roles.role1.name"))

	db = OpenDB("foo")

	// Watch the file and get a callback on change. The callback can do whatever,
	// like re-load the configuration.
	// File provider always returns a nil `event`.

	f.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("watch error: %v", err)
			return
		}

		// Throw away the old config and load a fresh copy.
		log.Println("config changed. Reloading ...")
		k = koanf.New(".")
		k.Load(f, yaml.Parser())
		k.Print()
	})

	log.Println("waiting forever. Try making a change to drawbridge.yaml to live reload")
	<-make(chan bool)
}

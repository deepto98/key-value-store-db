package main

import (
	"flag"
	"log"

	bolt "go.etcd.io/bbolt"
)

var (
	dbLocation = flag.String("db-location", "", "Path to bolt db")
)

func parseFlags() {
	flag.Parse()
	if *dbLocation == "" {
		log.Fatal("Missing db-location")
	}
}

func main() {
	parseFlags()

	db, err := bolt.Open(*dbLocation, 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

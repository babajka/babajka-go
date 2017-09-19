package main

import (
	"flag"
	"log"

	"github.com/go-xorm/xorm"
)

var (
	dev bool

	db           *xorm.Engine
	dbConnection string
)

func parseFlags() {
	flag.StringVar(&dbConnection, "db_connection", "", "to describe.")
	flag.BoolVar(&dev, "dev", false, "To run in dev mode.")
	flag.Parse()

	log.Printf("SERVER CONFIGURATION: ===========================\n")
	log.Printf("--dev = %v", dev)
	log.Printf("--db_connection = %v", dbConnection)
	log.Printf("=================================================\n")
}

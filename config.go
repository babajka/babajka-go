package main

import (
	"flag"
	"fmt"

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

	fmt.Println("Dev is on:", dev)
}

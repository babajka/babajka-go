package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func initDB() error {
	log.Println("initializing db...")
	if dbConnection == "" {
		dbConnection = os.Getenv("DATABASE_URL")
	}
	if dbConnection == "" {
		return errors.New("no DB connection rules provided")
	}
	log.Println("DB connection rules:", dbConnection)

	var err error
	db, err = xorm.NewEngine("postgres", dbConnection)
	if err != nil {
		return err
	}
	if dev {
		db.ShowSQL(true)
		if err := populateDB(); err != nil {
			return err
		}
	}

	return nil
}

func populateDB() error {
	has, err := db.Exist(&Article{})
	if err != nil {
		return fmt.Errorf("exist check failed: %v", err)
	}
	if has {
		log.Println("'Article' table already exists. Skipping population")
		return nil
	}

	db.CreateTables(&Article{})

	_, err = db.Insert(articles[1], articles[2], articles[3])
	if err != nil {
		return fmt.Errorf("insert error: %v", err)
	}

	c, err := db.Count(&Article{})
	if err != nil {
		return fmt.Errorf("count error: %v", err)
	}
	log.Printf("Number of articles populated: %d\n", c)

	return nil
}

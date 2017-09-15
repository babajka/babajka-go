package main

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func populateDB() {
	has, err := db.Exist(&Article{})
	if err != nil {
		fmt.Println("exist check failed: ", err)
	}
	if has {
		fmt.Println("article table already exists")
		return
	}

	db.CreateTables(&Article{})

	_, err = db.Insert(articles[1], articles[2], articles[3])
	if err != nil {
		fmt.Println("insert error: ", err)
	}

	c, err := db.Count(&Article{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of articles: ", c)

	tables, err := db.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, table := range tables {
		fmt.Println(table.Name)
	}
}

func initDB() {
	fmt.Println("initializing db...")
	if dbConnection == "" {
		dbConnection = os.Getenv("DATABASE_URL")
	}
	fmt.Println("DB CONNECTION RULES: ", dbConnection)

	var err error
	db, err = xorm.NewEngine("postgres", dbConnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	if dev {
		db.ShowSQL(true)
		populateDB()
	}
}

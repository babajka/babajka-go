package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var (
	dev bool

	db           *xorm.Engine
	dbConnection string
)

func defineStaticRoutes() {
	for _, r := range []string{"/styles/", "/images/", "/fonts/"} {
		http.Handle(r, http.StripPrefix(r, http.FileServer(http.Dir("./static"+r))))
	}
}

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

func parseFlags() {
	flag.StringVar(&dbConnection, "db_connection", "", "to describe.")

	devP := flag.Bool("dev", false, "To run in dev mode.")
	flag.Parse()
	dev = *devP
	fmt.Println("Dev is on:", dev)
}

func main() {
	fmt.Println("initializing server...")

	parseFlags()

	initDB()

	defineStaticRoutes()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/articles/", handleArticle)
	http.HandleFunc("/login/", handleLogin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("listening on port :%s\n", port)
	http.ListenAndServe(":"+port, nil)
	db.Close()
}

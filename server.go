package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func defineStaticRoutes() {
	for _, r := range []string{"/styles/", "/images/", "/fonts/"} {
		http.Handle(r, http.StripPrefix(r, http.FileServer(http.Dir("./static"+r))))
	}
}

func initDB() {
	_, err := sql.Open("postgres", "user=uladbohdan dbname=wir-prod sslmode=verify-full")
	if err != nil {
		fmt.Println("DB ERR", err)
		return
	}
}

func main() {
	fmt.Println("initializing server...")

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
}

package main

import (
	"log"
	"net/http"
	"os"
)

func defineStaticRoutes() {
	for _, r := range []string{"/styles/", "/images/", "/fonts/"} {
		http.Handle(r, http.StripPrefix(r, http.FileServer(http.Dir("./static"+r))))
	}
}

func main() {
	log.Println("initializing server...")
	parseFlags()
	if err := initDB(); err != nil {
		log.Println("failed to init db: ", err)
		return
	}
	defineStaticRoutes()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/articles/", handleArticle)
	http.HandleFunc("/login/", handleLogin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on port :%s\n", port)
	http.ListenAndServe(":"+port, nil)
	db.Close()
}

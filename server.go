package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handleArticle(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	aid, _ := strconv.Atoi(ss[2])
	if aid != 1 && aid != 2 {
		aid = 1
	}
	article := articles[aid]

	articleWrapper := ArticlePage{
		Article:  article,
		Featured: []int{1, 2, 3},
	}

	fmt.Println("requesting article", aid)
	tmpl, _ := template.ParseFiles("./templates/article.html")
	tmpl.Execute(w, articleWrapper)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html")
	tmpl.Execute(w, articles[1])
}

func defineStaticRoutes() {
	for _, r := range []string{"/styles/", "/images/", "/fonts/"} {
		http.Handle(r, http.StripPrefix(r, http.FileServer(http.Dir("."+r))))
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

	http.Handle("/", http.FileServer(http.Dir("./templates")))

	defineStaticRoutes()

	http.HandleFunc("/articles/", handleArticle)
	http.HandleFunc("/login/", handleLogin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("listening on port :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

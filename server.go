package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func handleArticle(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	article := articles[1]
	aid, _ := strconv.Atoi(ss[2])
	article.ID = aid
	fmt.Println("requesting article", aid)
	tmpl, _ := template.ParseFiles("./templates/article.html")
	tmpl.Execute(w, article)
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

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

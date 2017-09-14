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
	article := articles[aid]

	articleWrapper := ArticlePage{
		Article:  article,
		Featured: []Article{articles[1], articles[2], articles[3]},
	}

	fmt.Println("requesting article", aid)
	tmpl, _ := template.ParseFiles("./templates/article.html", "./templates/navbar.html")

	tmpl.Execute(w, articleWrapper)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html", "./templates/navbar.html")
	tmpl.Execute(w, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/index.html", "./templates/navbar.html")

	indexPage := &IndexPage{
		Top3: []Article{articles[1], articles[2], articles[3]},
	}

	tmpl.Execute(w, indexPage)
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

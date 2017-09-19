package main

import (
	"html/template"
	"net/http"
)

// IndexPage represents Main page.
type IndexPage struct {
	Top3                      []Article
	Filler1, Filler2, Filler3 Article
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(templateList("index/index", "navbar", "footer", "index/regular-tile", "header")...)

	indexPage := &IndexPage{
		Top3:    []Article{getArticleFromDB("1"), getArticleFromDB("2"), getArticleFromDB("3")},
		Filler1: getArticleFromDB("1"),
		Filler2: getArticleFromDB("2"),
		Filler3: getArticleFromDB("3"),
	}

	tmpl.Execute(w, indexPage)
}

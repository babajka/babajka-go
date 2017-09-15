package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Article model.
type Article struct {
	ID        int
	Title     string
	Text      string `xorm:"text"`
	Author    string
	Hashtags  []string
	Photo     string
	IsSpecial bool
}

// ArticlePage is a wrapper for Article to render the page.
type ArticlePage struct {
	Article  Article
	Featured []Article
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	// aid, _ := strconv.Atoi(ss[2])
	var article Article
	has, err := db.Where("i_d = ?", ss[2]).Get(&article)
	if err != nil {
		fmt.Println("article error: ", err)
	}
	if !has {
		fmt.Println("article with id ", ss[2], "not found")
	}

	articleWrapper := ArticlePage{
		Article:  article,
		Featured: []Article{articles[1], articles[2], articles[3]},
	}

	fmt.Println("requesting article", ss[2])
	tmpl, _ := template.ParseFiles(templateList("article/base", "navbar", "footer", "article/featured-tile", "header")...)

	tmpl.Execute(w, articleWrapper)
}

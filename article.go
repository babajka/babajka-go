package main

import (
	"html/template"
	"log"
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
	article := getArticleFromDB(ss[2])

	articleWrapper := ArticlePage{
		Article:  article,
		Featured: []Article{getArticleFromDB("1"), getArticleFromDB("2"), getArticleFromDB("3")},
	}

	tmpl, _ := template.ParseFiles(templateList("article/base", "navbar", "footer", "article/featured-tile", "header")...)

	tmpl.Execute(w, articleWrapper)
}

func getArticleFromDB(id string) Article {
	var article Article
	has, err := db.Where("i_d = ?", id).Get(&article)
	if err != nil {
		log.Println("article error:", err)
	}
	if !has {
		log.Println("article with id ", id, "not found")
	}
	return article
}

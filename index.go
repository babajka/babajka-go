package main

import (
	"html/template"
	"net/http"
)

// IndexPage ololo
type IndexPage struct {
	Top3 []Article
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/index.html", "./templates/navbar.html", "./templates/footer.html")

	indexPage := &IndexPage{
		Top3: []Article{articles[1], articles[2], articles[3]},
	}

	tmpl.Execute(w, indexPage)
}

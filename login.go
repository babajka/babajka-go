package main

import (
	"html/template"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(templateList("login/base", "navbar", "footer", "header")...)
	tmpl.Execute(w, nil)
}

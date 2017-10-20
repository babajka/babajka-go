package main

import (
	"html/template"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(templateList("login")...)
	tmpl.Execute(w, nil)
}

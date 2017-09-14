package main

import (
	"html/template"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html", "./templates/navbar.html", "./templates/footer.html")
	tmpl.Execute(w, nil)
}

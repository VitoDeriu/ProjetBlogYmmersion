package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Categorie(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Query().Get("type"))
	temps.Temp.ExecuteTemplate(w, "categorie", nil)
}

package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Ajout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Article data.ArtStruct
		Article.Name = r.FormValue("titre")
		Article.Content = r.FormValue("contenu")
		Article.Auteur = r.FormValue("author")
		AddArticle(Article, true)
	} else if r.Method == "GET" {
		temps.Temp.ExecuteTemplate(w, "ajout", nil)
	}
}

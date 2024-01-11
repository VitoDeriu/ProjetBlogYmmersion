package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
	"strings"
)

func Categorie(w http.ResponseWriter, r *http.Request) {
	data := SearchCategory(r.URL.Query().Get("type"))
	temps.Temp.ExecuteTemplate(w, "categorie", data)
}

// Fonction de recherche par catégorie
func SearchCategory(q string) []data.ArtStruct {
	var pertinentArticles []data.ArtStruct
	for _, article := range Articles {
		if strings.Contains(strings.ToLower(article.Category), strings.ToLower(q)) {
			pertinentArticles = append(pertinentArticles, article)
		}
	}
	return pertinentArticles
}

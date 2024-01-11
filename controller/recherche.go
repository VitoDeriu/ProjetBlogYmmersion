package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
	"strings"
)

func RechercheTemp(w http.ResponseWriter, r *http.Request) {
	data := SearchCategory(r.URL.Query().Get("type"))
	temps.Temp.ExecuteTemplate(w, "recherche", data)
}

// Fonction de recherche par Nom
func SearchName(q string) []data.ArtStruct {
	var pertinentArticles []data.ArtStruct
	for _, article := range Articles {
		if strings.Contains(strings.ToLower(article.Name), strings.ToLower(q)) {
			pertinentArticles = append(pertinentArticles, article)
			
		} else {
			
		}
	} 
	return pertinentArticles
}

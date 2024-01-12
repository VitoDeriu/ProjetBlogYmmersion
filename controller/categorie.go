package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"fmt"
	"net/http"
	"strings"
)

func Categorie(w http.ResponseWriter, r *http.Request) {
	dataa := SearchCategory(r.URL.Query().Get("type"))
	type test struct {
		Data []data.ArtStruct
		Type string
	}
	toSend := test{Data: dataa, Type: r.URL.Query().Get("type")}
	fmt.Println(r.URL.Query().Get("type"))
	//fmt.Println(toSend)
	//fmt.Println(dataa)
	temps.Temp.ExecuteTemplate(w, "categorie", toSend)
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

package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
	"strings"
)

func ArticleTemp(w http.ResponseWriter, r *http.Request) {
	data := SearchId(r.URL.Query().Get("id"))
	temps.Temp.ExecuteTemplate(w, "article", data)
}

func SearchId(q string) []data.ArtStruct {
	var pertinentArticles []data.ArtStruct
	for _, article := range Articles {
		if strings.Contains(strings.ToLower(article.Id), strings.ToLower(q)) {
			pertinentArticles = append(pertinentArticles, article)
		}
	}
	return pertinentArticles
}

package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
	"strconv"
)

func ArticleTemp(w http.ResponseWriter, r *http.Request) {
	/*data := SearchId(r.URL.Query().Get("type"))
	fmt.Println(data)
	fmt.Println(r.URL.Query().Get("type"))*/
	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)
	temps.Temp.ExecuteTemplate(w, "article", data)
}

// recherche par ID
func SearchId(id int) []data.ArtStruct {
	var pertinentArticles []data.ArtStruct
	for _, article := range Articles {
		if article.Id == id {
			pertinentArticles = append(pertinentArticles, article)
		}
	}
	return pertinentArticles
}

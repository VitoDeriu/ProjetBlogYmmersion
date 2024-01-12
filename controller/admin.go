package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := GetAllArticles()
		temps.Temp.ExecuteTemplate(w, "admin", data)
	} else if r.Method == "POST" {
		_ = r.ParseForm()
		if r.PostForm == nil {
			temps.Temp.ExecuteTemplate(w, "error404", nil)
			return
		}
		ArtId, _ := strconv.Atoi(r.PostForm.Get("Suppr"))
		if ArticleExist(int(ArtId)) {
			RemoveArticle(int(ArtId), true)
			data := GetAllArticles()
			temps.Temp.ExecuteTemplate(w, "admin", data)
		} else {
			temps.Temp.ExecuteTemplate(w, "error404", nil)
			return
		}
	}
}

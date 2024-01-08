package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func ArticleTemp (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "article", nil)
}
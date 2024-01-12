package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Admin (w http.ResponseWriter, r *http.Request){
data := GetAllArticles()

	temps.Temp.ExecuteTemplate(w, "admin", data)
}
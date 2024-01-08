package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Categorie (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "categorie", nil)
}
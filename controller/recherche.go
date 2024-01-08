package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func RechercheTemp (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "recherche", nil)
}
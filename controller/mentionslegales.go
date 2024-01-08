package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func MentionsLegales (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "legal", nil)
}
package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Explications (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "explications", nil)
}
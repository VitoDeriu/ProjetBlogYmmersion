package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Error404 (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "error404", nil)
}
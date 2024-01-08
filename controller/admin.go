package controller

import (
	"ProjetBlogYmmersion/temps"
	"net/http"
)

func Admin (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "admin", nil)
}
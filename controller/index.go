package controller

import (
	InitTemp "ProjetBlogYmmersion/temps"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "home", nil)
}

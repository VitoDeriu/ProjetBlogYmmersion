package controller

import (
	InitTemp "ProjetBlogYmmersion/temps"
	"encoding/json"
	"net/http"
)

var article ArtStruct

func GetDataFromJson() {

	os.ReadFile("ProjetBlogYmmersion/data")

	json.Unmarshal(article, &article)
}

func Home(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "home", nil)
}

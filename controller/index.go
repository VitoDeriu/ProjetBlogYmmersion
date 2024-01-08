package controller

import (
	"ProjetBlogYmmersion/data"
	InitTemp "ProjetBlogYmmersion/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var Articles []data.ArtStruct

func GetDataFromJson() {

	data, err := os.ReadFile("data/article.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Articles)
}

func Home(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	fmt.Println(Articles)
	InitTemp.Temp.ExecuteTemplate(w, "home", Articles)
}

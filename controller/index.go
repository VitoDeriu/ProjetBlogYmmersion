package controller

import (
	"ProjetBlogYmmersion/data"
	InitTemp "ProjetBlogYmmersion/temps"
	"encoding/json"
	"fmt"
	"math/rand"
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

	var dato []data.ArtStruct
	GetDataFromJson()
	dato = append(dato, Articles[1])
	//fmt.Println(Articles)
	//fmt.Println(dato)
	RandomArticle()
	RandomArticle()
	RandomArticle()
	RandomArticle()
	fmt.Println("---------------")
	InitTemp.Temp.ExecuteTemplate(w, "article", dato)
}

func RandomArticle() {
	var randomart []data.ArtStruct
	randomIndex := rand.Intn(len(Articles))

	for _, i := range randomart {
		if Articles[randomIndex] == i {
			//RandomArticle()
		} else {
			randomart = append(randomart, Articles[randomIndex])
		}
	}
	fmt.Println(randomart)
}

package controller

import (
	"ProjetBlogYmmersion/data"
	InitTemp "ProjetBlogYmmersion/temps"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"slices"
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
	//RandomArticle(4)
	fmt.Println("---------------")
	InitTemp.Temp.ExecuteTemplate(w, "home", RandomArticle(4))
}

func RandomArticle(nbrArt int) []data.ArtStruct {

	var randomart []data.ArtStruct

	for len(randomart) != nbrArt {
		randomIndex := rand.Intn(len(Articles) - 1)
		if !slices.Contains(randomart, Articles[randomIndex]) {
			randomart = append(randomart, Articles[randomIndex])
		}
	}
	fmt.Println(randomart)
	return randomart
}

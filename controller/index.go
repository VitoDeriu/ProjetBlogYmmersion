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

// mainfunc pour la page home
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/flavicon.ico" {
		InitTemp.Temp.ExecuteTemplate(w, "error404", nil)
		return
	}
	GetDataFromJson()
	InitTemp.Temp.ExecuteTemplate(w, "home", RandomArticle(4))

}

// récupération des données du json et envoie dans la structure ArtStruct
func GetDataFromJson() {
	data, err := os.ReadFile("data/article.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Articles) //passage en json vers la struct
}

// fonction pour récupèrer un nombre d'article aléatoire
func RandomArticle(nbrArt int) []data.ArtStruct {
	var randomart []data.ArtStruct
	for len(randomart) != nbrArt {
		randomIndex := rand.Intn(len(Articles) - 1)
		if !slices.Contains(randomart, Articles[randomIndex]) {
			randomart = append(randomart, Articles[randomIndex])
		}
	}
	return randomart
}

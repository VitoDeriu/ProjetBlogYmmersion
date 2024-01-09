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

// ajout des données de Articles dans notre fichier JSON
func SetDataToJson() {
	data, err := json.Marshal(Articles)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/article.json", data, 0644)
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
	fmt.Println(randomart)
	return randomart
}

// fonction pour ajouter un article à notre tableau et potentiellement au json
func AddArticle(article data.ArtStruct, save bool) {
	GetDataFromJson()
	Articles = append(Articles, article)
	if save {
		SetDataToJson()
	}
}

// fonction pour supprimer un article de notre tableau et potentiellement du json
func RemoveArticle(index int, save bool) {
	GetDataFromJson()
	for i := 0; i < len(Articles); i++ {
		if i == index {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
	if save {
		SetDataToJson()
	}
}

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
var Dato []data.ArtStruct

// récupération des données du json et envoie dans la structure ArtStruct
func GetDataFromJson() {

	data, err := os.ReadFile("data/article.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Articles) //passage en go vers la struct
}

// faut faire la fonction qui récup aléatoirement les articles, qui les balances dans une variable qu'on récup pour afficher dans le html
func GetRandomArticles() []data.ArtStruct {

	//prendre la length de Articles, prendre 4 nombres aléatoires compris entre 0 et la length et mettre ces nombres commes les index
	//pour les envoyer dans dato 

	Dato = append(Dato, Articles[0], Articles[1], Articles[2], Articles[3])

	fmt.Println(Dato)
	return Dato
}

// mainfunc pour la page home
func Home(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	GetRandomArticles()

	// fmt.Println(Articles)
	InitTemp.Temp.ExecuteTemplate(w, "home", Dato)
}

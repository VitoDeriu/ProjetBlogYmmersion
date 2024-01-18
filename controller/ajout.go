package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Ajout(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "ajout", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Récup du fichier img
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//Créa de l'img dans le dossier /assets/img
	path, _ := os.Getwd()
	filePath := filepath.Join(path, "/assets/img/", handler.Filename)
	outFile, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}

	//pour fermer le fichier
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	// fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	//si j'ai bien compris, on stocke le fichier img dans file, le path du fichier dans filepath et on enregistre le tout (donc l'image) dans le projet dans le dossier img avec io.copy.

	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Article data.ArtStruct
		Article.Name = r.FormValue("title")
		Article.Content = r.FormValue("content")
		Article.Auteur = r.FormValue("author")
		Article.Category = r.FormValue("category")
		Article.Preview = r.FormValue("preview")
		Article.Date = GetCurrentTime()
		Article.Img = handler.Filename
		Article.Id = GetArticleId()
		AddArticle(Article, true)
	}

	temps.Temp.ExecuteTemplate(w, "ajout", nil) //pas sur de celle là, elle réaffiche la page sans aucun feedback
}

// fonction pour ajouter un article à notre tableau et potentiellement au json
func AddArticle(article data.ArtStruct, save bool) {
	GetDataFromJson()
	Articles = append(Articles, article)
	if save {
		SetDataToJson()
	}
}

// Fonction qui récup la date et l'heure
func GetCurrentTime() string {

	currentTime := time.Now()
	timeFormatted := currentTime.Format("02-January-2006, 15:04")

	return timeFormatted
}

// Fonction qui récup la longueur du tableau d'Articles +1 pour un nouvel article
func GetArticleId() int {
	fmt.Println(len(Articles))	
	id := len(Articles) + 1
	fmt.Println(id)
	return id
}

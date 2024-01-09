package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Ajout (w http.ResponseWriter, r *http.Request){
	temps.Temp.ExecuteTemplate(w, "ajout", nil)
}

/*var filePath string
func Ajout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Article data.ArtStruct
		Article.Name = r.FormValue("title")
		Article.Content = r.FormValue("content")
		Article.Auteur = r.FormValue("author")
		Article.Category = r.FormValue("category")
		Article.Date = GetCurrentTime()
		Article.Img = filePath
		fmt.Println(filePath)
		AddArticle(Article, true)
	} else if r.Method == "GET" {
		temps.Temp.ExecuteTemplate(w, "ajout", nil)
	}
}*/

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

	filePath := filepath.Join(/*"assets", "img",*/ handler.Filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}
	
	defer outFile.Close()



	_, err = io.Copy(outFile, file)
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	//time.Sleep(3 * time.Second)
	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Article data.ArtStruct
		Article.Name = r.FormValue("title")
		Article.Content = r.FormValue("content")
		Article.Auteur = r.FormValue("author")
		Article.Category = r.FormValue("category")
		Article.Date = GetCurrentTime()
		Article.Img = filePath
		fmt.Println(filePath)
		AddArticle(Article, true)
	} else if r.Method == "GET" {
		temps.Temp.ExecuteTemplate(w, "ajout", nil)
	}
}


package controller

import (
	"ProjetBlogYmmersion/data"
	"ProjetBlogYmmersion/temps"
	"net/http"
	"path/filepath"
	"fmt"
	"os"
	"io"
)

func Ajout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Article data.ArtStruct
		Article.Name = r.FormValue("titre")
		Article.Content = r.FormValue("contenu")
		Article.Auteur = r.FormValue("author")
		Article.Category = r.FormValue("category")
		Article.Date = GetCurrentTime()
		//Article.Img = 
		AddArticle(Article, true)
	} else if r.Method == "GET" {
		temps.Temp.ExecuteTemplate(w, "ajout", nil)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
        return
    }
	defer file.Close()

	filePath := filepath.Join("./","assets", "img", handler.Filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file")
        fmt.Println(err)
        return
    }

    defer outFile.Close()
	
	_, err = io.Copy(outFile,file)
	if err != nil {
        fmt.Println("Error writing the file")
        fmt.Println(err)
        return
    }
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}



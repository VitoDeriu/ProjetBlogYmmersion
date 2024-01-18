package routeur

import (
	"ProjetBlogYmmersion/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/admin", controller.Admin)
	http.HandleFunc("/ajout", controller.Ajout)
	http.HandleFunc("/article", controller.ArticleTemp)
	http.HandleFunc("/categorie", controller.Categorie)
	http.HandleFunc("/explication", controller.Explications)
	http.HandleFunc("/mentionslegales", controller.MentionsLegales)
	http.HandleFunc("/rechercher", controller.RechercheTemp)
	http.HandleFunc("/upload", controller.UploadFile)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8090/) - Server started on port:8090")
	http.ListenAndServe("localhost:8090", nil)

}

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
	http.HandleFunc("/mentions-legales", controller.MentionsLegales)
	http.HandleFunc("/rechercher", controller.RechercheTemp)

	//ne pas oublier de faire la route error 404, je sais plus comment elle se fait y'a un truc particulier dans la gestion d'erreur a mettre.

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)

}

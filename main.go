package main

import (
	"ProjetBlogYmmersion/controller"
	"ProjetBlogYmmersion/routeur"
	"ProjetBlogYmmersion/temps"
)

func main() {
	controller.GetDataFromJson()
	temps.InitTemplate()
	routeur.InitServ()
}

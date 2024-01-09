package main

import (
	"ProjetBlogYmmersion/routeur"
	"ProjetBlogYmmersion/temps"
)

func main() {
	temps.InitTemplate()
	routeur.InitServ()
}

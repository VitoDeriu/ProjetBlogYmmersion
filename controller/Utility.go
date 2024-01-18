	package controller

import (
	"ProjetBlogYmmersion/data"
	"encoding/json"
	"fmt"
	"os"
)



// ajout des données de Articles dans notre fichier JSON
func SetDataToJson() {
	data, err := json.Marshal(Articles)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/article.json", data, 0644)
}

// fonction pour récupèrer les articles d'une catégorie précise
func GetArticles(category string) []data.ArtStruct {
	GetDataFromJson()
	var articles []data.ArtStruct
	for _, article := range Articles {
		if article.Category == category {
			articles = append(articles, article)
		}
	}
	return articles
}

// fonction pour récupèrer tous les articles toutes catégories confondues
func GetAllArticles() []data.ArtStruct {
	GetDataFromJson()
	return Articles
}

func ArticleExist(id int) bool {
	GetDataFromJson()
	for _, art := range Articles {
		if art.Id == id {
			return true
		}
	}
	return false
}

// fonction pour supprimer un article de notre tableau et potentiellement du json
func RemoveArticle(index int, save bool) {
	GetDataFromJson()
	var NewArt []data.ArtStruct
	for _, art := range Articles {
		if art.Id != index {
			NewArt = append(NewArt, art)
		}
	}
	Articles = NewArt
	if save {
		SetDataToJson()
	}
}
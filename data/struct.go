package data

type ArtStruct struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Img      string `json:"img"`
	Auteur   string `json:"auteur"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Preview  string `json:"preview"`
}

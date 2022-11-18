package models

type Imdb struct {
	Rating interface{} `json:"rating" bson:"rating"`
	Votes  interface{} `json:"votes" bson:"votes"`
	Id     interface{} `json:"id" bson:"id"`
}

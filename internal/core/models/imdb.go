package models

type Imdb struct {
	Rating interface{} `bson:"rating"`
	Votes  interface{} `bson:"votes"`
	Id     interface{} `bson:"id"`
}

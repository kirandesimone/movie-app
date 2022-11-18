package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Plot         string             `json:"plot" bson:"plot,omitempty"`
	Genres       []string           `json:"genres" bson:"genres,omitempty"`
	Runtime      int32              `json:"runtime" bson:"runtime,omitempty"`
	Rated        string             `json:"rated" bson:"rated,omitempty"`
	Title        string             `json:"title" bson:"title,omitempty"`
	Poster       string             `json:"poster" bson:"poster,omitempty"`
	Countries    []string           `json:"countries" bson:"countries,omitempty"`
	FullPlot     string             `json:"fullplot" bson:"fullplot,omitempty"`
	CommentCount int32              `json:"num_mflix_comments" bson:"num_mflix_comments,omitempty"`
	Cast         []string           `json:"cast" bson:"cast,omitempty"`
	Directors    []string           `json:"directors" bson:"directors,omitempty"`
	Writers      []string           `json:"writers" bson:"writers,omitempty"`
	Awards       *Award             `json:"awards" bson:"awards,omitempty"`
	Updated      string             `json:"lastupdated" bson:"lastupdated,omitempty"`
	Year         interface{}        `json:"year" bson:"year,omitempty"`
	Imdb         *Imdb              `json:"imdb" bson:"imdb,omitempty"`
	Type         string             `json:"type" bson:"type,omitempty"`
	Tomatoes     *Tomato            `json:"tomatoes" bson:"tomatoes,omitempty"`
}

func NewMovie(plot string, genres []string, runtime int32, rated string, title string) {

}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id           primitive.ObjectID `bson:"_id"`
	Plot         string             `bson:"plot,omitempty"`
	Genres       []string           `bson:"genres,omitempty"`
	Runtime      int32              `bson:"runtime,omitempty"`
	Rated        string             `bson:"rated,omitempty"`
	Title        string             `bson:"title,omitempty"`
	Poster       string             `bson:"poster,omitempty"`
	Countries    []string           `bson:"countries,omitempty"`
	FullPlot     string             `bson:"fullplot,omitempty"`
	CommentCount int32              `bson:"num_mflix_comments,omitempty"`
	Cast         []string           `bson:"cast,omitempty"`
	Directors    []string           `bson:"directors,omitempty"`
	Writers      []string           `bson:"writers,omitempty"`
	Awards       *Award             `bson:"awards,omitempty"`
	Updated      string             `bson:"lastupdated,omitempty"`
	Year         interface{}        `bson:"year,omitempty"`
	Imdb         *Imdb              `bson:"imdb,omitempty"`
	Type         string             `bson:"type,omitempty"`
	Tomatoes     *Tomato            `bson:"tomatoes,omitempty"`
}

func NewMovie(plot string, genres []string, runtime int32, rated string, title string) {

}

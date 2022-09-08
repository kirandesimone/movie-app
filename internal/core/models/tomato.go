package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tomato struct {
	Viewers    *Viewer            `bson:"viewer,omitempty"`
	Fresh      int32              `bson:"fresh,omitempty"`
	Critics    *Critic            `bson:"critic,omitempty"`
	Rotten     int32              `bson:"rotten,omitempty"`
	Updated    primitive.DateTime `bson:"lastUpdated,omitempty"`
	Dvd        primitive.DateTime `bson:"dvd,omitempty"`
	Consensus  string             `bson:"consensus,omitempty"`
	Production string             `bson:"production,omitempty"`
}

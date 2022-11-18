package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tomato struct {
	Viewers    *Viewer            `json:"viewer" bson:"viewer,omitempty"`
	Fresh      int32              `json:"fresh" bson:"fresh,omitempty"`
	Critics    *Critic            `json:"critic" bson:"critic,omitempty"`
	Rotten     int32              `json:"rotten" bson:"rotten,omitempty"`
	Updated    primitive.DateTime `json:"lastUpdated" bson:"lastUpdated,omitempty"`
	Dvd        primitive.DateTime `json:"dvd" bson:"dvd,omitempty"`
	Consensus  string             `json:"consensus" bson:"consensus,omitempty"`
	Production string             `json:"production" bson:"production,omitempty"`
}

package models

type Award struct {
	Wins        int32  `json:"wins" bson:"wins"`
	Nominations int32  `json:"nominations" bson:"nominations"`
	Text        string `json:"text" bson:"text"`
}

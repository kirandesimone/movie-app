package models

type Award struct {
	Wins        int32  `bson:"wins"`
	Nominations int32  `bson:"nominations"`
	Text        string `bson:"text"`
}

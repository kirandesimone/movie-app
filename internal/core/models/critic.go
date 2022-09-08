package models

type Critic struct {
	Rating     float64 `bson:"rating"`
	NumReviews int32   `bson:"numReviews"`
	Meter      int32   `bson:"meter"`
}

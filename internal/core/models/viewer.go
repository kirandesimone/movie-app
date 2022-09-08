package models

type Viewer struct {
	Rating     float64 `bson:"rating"`
	NumReviews int32   `bson:"numReviews"`
	Meter      int32   `bson:"meter"`
}

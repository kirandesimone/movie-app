package models

type Viewer struct {
	Rating     float64 `json:"rating" bson:"rating"`
	NumReviews int32   `json:"numReviews" bson:"numReviews"`
	Meter      int32   `json:"meter" bson:"meter"`
}

package model

// Reply holds properties related to reply entity in db
type Reply struct {
	ID            string  `bson:"_id"`
	Text          string  `bson:"text"`
	MinConfidence float64 `bson:"min_confidence"`
	Intent        string  `bson:"intent"`
}

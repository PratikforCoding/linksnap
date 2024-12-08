package models

import "time"

// URL is the model for the URL collection
type URL struct {
	ID			string		`bson:"_id,omitempty"`
	ShortCode 	string		`bson:"shortCode" validate:"required"`
	LongURL 	string		`bson:"longUrl" validate:"required,url"`
	CustomAlias	string 		`bson:"customAlias, omitempty"`
	CreatedAt 	time.Time 	`bson:"createdAt"`
	Clicks 		int			`bson:"clicks"`
}
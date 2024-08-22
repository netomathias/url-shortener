package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlShortener struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	OriginalUrl string `json:"original_url" bson:"original_url"`
	ShortenedUrl string `json:"shortened_url" bson:"shortened_url"`
	Alias string `json:"alias" bson:"alias"`
	Clicks int `json:"clicks" bson:"clicks"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUrlShortenerModel() *UrlShortener {
	return &UrlShortener{
		Id: primitive.NewObjectID(),
		OriginalUrl: "",
		ShortenedUrl: "",
		Alias: "",
		Clicks: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
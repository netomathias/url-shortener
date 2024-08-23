package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlShortener struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	LongUrl string `json:"long_url" bson:"long_url"`
	ShortUrl string `json:"short_url" bson:"short_url"`
	Alias string `json:"alias" bson:"alias"`
	Clicks int `json:"clicks" bson:"clicks"`
	RaichuId string `json:"raichu_id" bson:"raichu_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUrlShortenerModel() *UrlShortener {
	return &UrlShortener{
		Id: primitive.NewObjectID(),
		LongUrl: "",
		ShortUrl: "",
		Alias: "",
		Clicks: 0,
		RaichuId: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
package repositories

import (
	"context"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlRepository interface {
	Save(request *models.UrlShortener) error
	FindByAlias(alias string) (models.UrlShortener, error)
	UpdateClicks(alias string) error
}

type UrlRepositoryImpl struct {
	Db *mongo.Database
}

func NewUrlRepositoryImpl(db *mongo.Database) UrlRepository {
	return &UrlRepositoryImpl{Db: db}
}

func (r *UrlRepositoryImpl) Save(request *models.UrlShortener) error {
	_, err := r.Db.Collection("url-shortener").InsertOne(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (r *UrlRepositoryImpl) FindByAlias(alias string) (models.UrlShortener, error) {
	var result models.UrlShortener
	err := r.Db.Collection("url-shortener").FindOne(context.Background(), bson.M{"alias": alias}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.UrlShortener{}, err
		}
	}

	return result, nil
}

func (r *UrlRepositoryImpl) UpdateClicks(alias string) error {
	result, err := r.FindByAlias(alias)
	if err != nil {
		return err
	}

	_, err = r.Db.Collection("url-shortener").UpdateOne(context.Background(), bson.M{"alias": alias}, bson.M{"$set": bson.M{"clicks": result.Clicks + 1}})
	if err != nil {
		return err
	}

	return nil
}
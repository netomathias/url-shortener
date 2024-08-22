package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RedisDB struct {
	Client *redis.Client
}

func NewRedisDB(config *RedisConfig) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.DB,  // use default DB
	}) 

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{Client: client}, nil
}
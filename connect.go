package main

import "github.com/redis/go-redis/v9"

func Connect(host, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	return client
}

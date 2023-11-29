package main

import (
	"context"
	"fmt"
	"time"
)

const redisHost = "localhost:6379"
const redisPwd = ""

func main() {
	// initialize redis connection
	dbRedis := Connect(redisHost, redisPwd)
	fmt.Println("redis client initialized")

	// storing data using set command
	key := "greeting"
	data := "Hi, I'm Muhammad Arif Sulaksono"
	ttl := time.Duration(3) * time.Second

	opSet := dbRedis.Set(context.Background(), key, data, ttl)
	if err := opSet.Err(); err != nil {
		fmt.Printf("failed to set data : %v", err)
		return
	}

	fmt.Println("set operation success")
}

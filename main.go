package main

import "fmt"

const redisHost = "localhost:6379"
const redisPwd = ""

func main() {
	// initialize redis connection
	dbRedis := Connect(redisHost, redisPwd)
	fmt.Println("redis client initialized on", dbRedis.String())
}

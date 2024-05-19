package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-gin-be-clean-arch/redis"
)

func NewRedisCache(env *Env) redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisUser := env.RedisUser
	redisPass := env.RedisPass
	redisHost := env.RedisHost
	redisPort := env.RedisPort
	redisDB := env.RedisDB

	// sample = redis://user:password@localhost:6379/0
	redisURI := fmt.Sprintf("redis://%s:%s@%s:%s/%d", redisUser, redisPass, redisHost, redisPort, redisDB)
	log.Println(redisURI)

	client, err := redis.NewClient(redisURI)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success connection to Redis")
	return client
}

func CloseRedisConnection(client redis.Client) {
	if client == nil {
		return
	}

	err := client.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to Redis closed.")
}

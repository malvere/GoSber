package redisstore

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func New() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("foo", val)
}

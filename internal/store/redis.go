package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()

func NewRedisCache(addr, pw string, db int) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Could not connect to Redis: %v", err))
	}
	
	return client
}


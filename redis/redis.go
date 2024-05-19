package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Ping(context.Context) (string, error)
	Close(context.Context) error
	Set(context.Context, string, interface{}, time.Duration) (string, error)
	Get(context.Context, string) (string, error)
	TTL(context.Context, string) (time.Duration, error)
}

type redisClient struct {
	cl *redis.Client
}

func NewClient(connection string) (Client, error) {
	opts, err := redis.ParseURL(connection)
	if err != nil {
		return &redisClient{}, err
	}
	c := redis.NewClient(opts)
	return &redisClient{cl: c}, nil
}

func (rc *redisClient) Ping(ctx context.Context) (string, error) {
	return rc.cl.Ping(ctx).Result()
}

func (rc *redisClient) Close(ctx context.Context) error {
	return rc.cl.Close()
}

func (rc *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return rc.cl.Set(ctx, key, value, expiration).Result()
}

func (rc *redisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.cl.Get(ctx, key).Result()
}

func (rc *redisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	return rc.cl.TTL(ctx, key).Result()
}

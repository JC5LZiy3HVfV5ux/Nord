package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	client *redis.Client
}

func newRedisClient(client *redis.Client) *redisClient {
	return &redisClient{
		client: client,
	}
}

func (r *redisClient) get(ctx context.Context, model interface{}, key string) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), model); err != nil {
		return err
	}

	return nil
}

func (r *redisClient) put(ctx context.Context, model interface{}, key string, expiration time.Duration) error {
	val, err := json.Marshal(model)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, key, val, expiration).Err(); err != nil {
		return err
	}

	return nil
}

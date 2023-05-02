package rediscache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func serialize(model any) ([]byte, error) {
	return json.Marshal(model)
}

func deserializee[T any](data []byte) (T, error) {
	var model T
	err := json.Unmarshal(data, &model)
	return model, err
}

func Set(ctx context.Context, rdb *redis.Client, key string, value any, ttl int) error {
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = rdb.Set(ctx, key, string(b), time.Duration(ttl)*time.Second).Result()
	return err
}

func Get[T any](ctx context.Context, rdb *redis.Client, key string) (T, error) {
	var row T
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return row, err
	}

	b, err := deserializee[T]([]byte(result))
	if err != nil {
		return row, err
	}
	return b, nil
}

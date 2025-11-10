package cache

import (
	"context"
	"go-app/pkg/errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisStore is a store for Redis
type redisStore struct {
	client *redis.Client
}

// NewRedisStore create cache instance with redis
func NewRedisStore(rd *redis.Client) Client {
	return &redisStore{
		client: rd,
	}
}

// Get value from key
func (rd redisStore) Get(ctx context.Context, k string) ([]byte, error) {
	obj, err := rd.client.Get(ctx, k).Bytes()

	switch {
	case errors.Is(err, redis.Nil):
		return []byte{}, errors.ErrRedisKeyNotFound.Trace()
	case err != nil:
		return []byte{}, errors.ErrRedisConnection.Wrap(err)
	default:
		return obj, nil
	}
}

// Set value by key and duration time
func (rd redisStore) Set(ctx context.Context, k string, v any, exp time.Duration) error {
	if err := rd.client.Set(ctx, k, v, exp).Err(); err != nil {
		return errors.ErrRedisConnection.Wrap(err)
	}

	return nil
}

// Del values keys
func (rd redisStore) Del(ctx context.Context, ks ...string) error {
	if _, err := rd.client.Del(ctx, ks...).Result(); err != nil {
		return errors.ErrRedisConnection.Wrap(err)
	}

	return nil
}

// FlushAll flush all data
func (rd redisStore) FlushAll(ctx context.Context) error {
	if err := rd.client.FlushAll(ctx).Err(); err != nil {
		return errors.ErrRedisConnection.Wrap(err)
	}

	return nil
}

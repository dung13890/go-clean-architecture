//go:generate mockgen -source=$GOFILE -destination=mock/cache_mock.go

package cache

import (
	"context"
	"time"
)

// Client is a interface for multiple store
type Client interface {
	Get(ctx context.Context, k string) ([]byte, error)
	Set(ctx context.Context, k string, v any, e time.Duration) error
	Del(ctx context.Context, ks ...string) error
	FlushAll(ctx context.Context) error
}

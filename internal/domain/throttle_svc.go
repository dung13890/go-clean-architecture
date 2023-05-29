//go:generate mockgen -source=$GOFILE -destination=mock/throttle_svc_mock.go
package domain

import (
	"context"
)

// ThrottleService is a struct that represent the throttle's service
type ThrottleService interface {
	Blocked(ctx context.Context, key, ip string) (bool, error)
	Incr(ctx context.Context, key, ip string) error
	Clear(ctx context.Context, key, ip string) error
}

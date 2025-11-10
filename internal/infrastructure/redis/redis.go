package redis

import (
	"fmt"
	"strconv"

	"go-app/internal/infrastructure/config"

	"github.com/redis/go-redis/v9"
)

// New setup Redis
func New(rd config.Redis) *redis.Client {
	uri := fmt.Sprintf("%s:%s",
		rd.Host,
		strconv.Itoa(rd.Port),
	)

	return redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: rd.Password, // no password set
		DB:       0,           // use default DB
	})
}

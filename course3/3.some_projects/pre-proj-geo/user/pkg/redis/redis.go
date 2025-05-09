package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Redis struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Client *redis.Client
}

func New(addr string, opts ...Option) (*Redis, error) {
	r := &Redis{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(r)
	}

	r.Client = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "",
		DB:          0,
		MaxRetries:  r.connAttempts,
		DialTimeout: r.connTimeout,
		PoolSize:    r.maxPoolSize,
	})

	log.Printf("Redis status: %s", r.Client.Ping().Val())

	return r, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}

package datastore

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

var ctx = context.Background()

func NewRedisClient(host map[string]string, pass string, db int) *redis.Ring {
	client := redis.NewRing(
		&redis.RingOptions{
			Addrs:        host,
			Password:     pass,
			DB:           db,
			DialTimeout:  time.Duration(30) * time.Second,
			WriteTimeout: time.Duration(30) * time.Second,
			ReadTimeout:  time.Duration(30) * time.Second,
		})

	client.AddHook(nrredis.NewHook(nil))

	if _, err := client.WithContext(ctx).Ping().Result(); err != nil {
		panic(err)
	}
	return client
}

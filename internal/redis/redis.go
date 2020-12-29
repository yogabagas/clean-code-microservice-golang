package redis

import (
	"context"
	"time"

	cache "github.com/go-redis/cache/v7"
	redis "github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type InternalRedisImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
}

type InternalRedis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, value interface{}) error
}

func NewRedisImpl(rdb *redis.Ring) InternalRedis {
	return &InternalRedisImpl{
		rdb: rdb,
		cache: &cache.Codec{
			Redis: rdb,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
	}
}

func (r *InternalRedisImpl) Set(ctx context.Context, key string, value interface{}) error {
	if err := r.cache.Set(&cache.Item{
		Ctx:        ctx,
		Key:        key,
		Object:     value,
		Expiration: time.Hour,
	}); err != nil {
		return err
	}
	return nil
}

func (r *InternalRedisImpl) Get(ctx context.Context, key string, value interface{}) error {
	if err := r.cache.GetContext(ctx, key, value); err != nil {
		return err
	}
	return nil
}

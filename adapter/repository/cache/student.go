package cache

import (
	"context"
	"my-github/clean-code-microservice-golang/usecase/student/repository"
	"time"

	"github.com/go-redis/cache/v7"
	redis "github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type StudentCacheRepository struct {
	redis *redis.Ring
	cache *cache.Codec
}

func NewRedisRepository(redis *redis.Ring) repository.StudentCacheRepository {
	return &StudentCacheRepository{
		redis: redis,
		cache: &cache.Codec{
			Redis: redis,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
			// LocalCache: cache.NewTinyLFU(1000, time.Minute),
		},
	}
}

func (s *StudentCacheRepository) WriteStudent(ctx context.Context, key string, value interface{}) error {
	if err := s.cache.Set(&cache.Item{
		Ctx:        ctx,
		Key:        key,
		Object:     value,
		Expiration: time.Hour,
	}); err != nil {
		return err
	}
	return nil

}

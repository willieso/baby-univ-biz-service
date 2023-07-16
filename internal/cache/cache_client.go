package cache

import (
	"context"

	"github.com/willieso/baby-univ-biz-service/internal/model"
	"github.com/willieso/baby-univ-biz-service/pkg/cache"
	"github.com/willieso/baby-univ-biz-service/pkg/encoding"
	"github.com/willieso/baby-univ-biz-service/pkg/redis"
)

func getCacheClient(ctx context.Context) cache.Cache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""
	client := cache.NewRedisCache(redis.RedisClient, cachePrefix, jsonEncoding, func() interface{} {
		return &model.UserBaseModel{}
	})

	return client
}

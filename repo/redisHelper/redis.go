package redisHelper

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"URLShortener/common/errorx"
	"URLShortener/log"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Stop() {
	if rdb != nil {
		_ = rdb.Close()
	}
}

func GetDB() *redis.Client {
	if rdb == nil {
		Init()
	}

	return rdb
}

func Get(ctx context.Context, key string) (val string, err *errorx.ServiceError) {
	var rdbErr error
	val, rdbErr = GetDB().Get(ctx, key).Result()
	if rdbErr != nil {
		if rdbErr == redis.Nil {
			return val, errorx.DataNotFoundError
		}

		log.Error(ctx, "%v", rdbErr)
		return val, err
	}

	return val, nil
}

func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (success bool, err *errorx.ServiceError) {
	var rdbErr error
	success, rdbErr = GetDB().SetNX(ctx, key, value, expiration).Result()
	if rdbErr != nil {
		log.Error(ctx, "%v", rdbErr)
		return false, errorx.InsertDataBaseFailedError
	}

	return success, nil
}

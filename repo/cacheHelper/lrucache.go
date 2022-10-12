package cacheHelper

import (
	"context"

	lru "github.com/hashicorp/golang-lru"

	"URLShortener/common/define"
	"URLShortener/log"
)

var cache *lru.Cache

func Init() {
	var err error
	cache, err = lru.New(define.CacheSize)
	if err != nil {
		log.Panic(context.Background(), "%v", err)
	}
}

func Add(ctx context.Context, key, value interface{}) {
	cache.Add(key, value)

	log.Info(ctx, "Cache_Add(%v) = %v", key, value)
}

func Get(ctx context.Context, key interface{}) (value interface{}, ok bool) {
	value, ok = cache.Get(key)

	log.Info(ctx, "Cache_Get(%v) = %v", key, value)

	return value, ok
}

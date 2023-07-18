package cache

import (
	"errors"
	"reflect"
	"sync"
)

type Cache interface {
	Type() string
	Set(key string, value interface{}, ttlSeconds int) error
	Get(key string, result interface{}) error
}

type CacheConfig struct {
	Type   string
	Option interface{}
}

var initOnce sync.Once

var defaultCache Cache

var ErrNotFound = errors.New("cache value of key is not found")

func Init(cacheType string, option interface{}) (err error) {
	initOnce.Do(func() {
		defaultCache, err = NewCache(cacheType, option)
	})
	return
}

func GetDefaultCache() Cache {
	return defaultCache
}

func NewCache(cacheType string, option interface{}) (cache Cache, err error) {
	switch cacheType {
	// TODO redis cache
	default:
		cache = newTtlMapCache()
	}
	if reflect.ValueOf(cache).IsNil() {
		cache = newTtlMapCache()
	}
	return
}

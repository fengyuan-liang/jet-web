package cache

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/mailgun/ttlmap"
)

type TtlMapCache struct {
	cache *ttlmap.TtlMap
}

var ttlMapCache *TtlMapCache
var initTtlMapOnce sync.Once

func newTtlMapCache() *TtlMapCache {
	initTtlMapOnce.Do(func() {
		cache, _ := ttlmap.NewConcurrent(65535)
		ttlMapCache = &TtlMapCache{
			cache: cache,
		}
	})
	return ttlMapCache
}

func (c *TtlMapCache) Set(key string, value interface{}, ttlSeconds int) error {
	return c.cache.Set(key, value, ttlSeconds)
}

func (c *TtlMapCache) Get(key string, result interface{}) error {
	val, ok := c.cache.Get(key)
	if !ok {
		return ErrNotFound
	}
	if result == nil {
		return fmt.Errorf("receiver result is nil")
	}
	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("result type must be pointer")
	}
	value := reflect.ValueOf(result).Elem()
	if !value.CanSet() {
		return fmt.Errorf("%s must cant get", value.String())
	}
	value.Set(reflect.ValueOf(val))

	return nil
}

func (c *TtlMapCache) Type() string {
	return "ttlmap"
}

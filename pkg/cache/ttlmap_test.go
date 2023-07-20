package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTtlMapCache_Set(t *testing.T) {
	ttl := newTtlMapCache()
	ttl.Set("name", "king", 3)
	name := "beego"
	err := ttl.Get("name", &name)
	assert.Nil(t, err)
	assert.Equal(t, "king", name)
	time.Sleep(time.Second * 3)
	err = ttl.Get("name", &name)
	// 3s过期
	assert.EqualError(t, err, "cache value of key is not found")
}

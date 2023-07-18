package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTtlMapCache_Set(t *testing.T) {
	ttl := newTtlMapCache()
	ttl.Set("name", "king", 3)
	name := "beego"
	err := ttl.Get("name", &name)
	assert.Nil(t, err)
	assert.Equal(t, "king", name)
}

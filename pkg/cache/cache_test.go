package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetDefaultCache(t *testing.T) {
	Init("", "")
	cache := GetDefaultCache()
	assert.Equal(t, "ttlmap", cache.Type())
}

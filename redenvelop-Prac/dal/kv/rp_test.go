package kv

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLPushRpAndLLenRpAndLPop(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// Set up test data
	key := "test_key"
	val := []string{"a", "b", "c"}

	// Clean up previous test data
	rdb.Del(ctx, getKey(key))

	// Test LPushRp
	err := LPushRp(ctx, key, val)
	assert.NoError(t, err)

	// Test LLenRp
	result, err := LLenRp(ctx, key)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(val), len(*result))

	// Test LPop
	for i := 0; i < len(val); i++ {
		result, err := LPop(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, int64(i), result)
	}

	// Clean up test data
	rdb.Del(ctx, getKey(key))
}

package util

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func GenerateTraceID() string {
	now := time.Now().UnixNano()

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	return fmt.Sprintf("%d-%04d", now, num)
}

func GetValueFromContext(ctx context.Context, key string) string {
	value := ""
	id := ctx.Value(key)
	if id != nil {
		value = id.(string)
	}
	return value
}

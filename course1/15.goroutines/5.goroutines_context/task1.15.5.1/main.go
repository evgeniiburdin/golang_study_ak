package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
}

func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {
	var cancel context.CancelFunc
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(contextDeadline))
	defer cancel()

	select {
	case <-ctx.Done():
		return "context deadline exceeded"
	case <-time.After(timeAfter):
		return "time after exceeded"
	}
}

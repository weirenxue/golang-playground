package main

import (
	"fmt"
	"time"
)

var tickerDuration = 1 * time.Microsecond

type Count struct {
	n        *int64
	duration *time.Duration
}

func Ptr[T any](n T) *T {
	return &n
}

func main() {
	count := Count{n: Ptr[int64](0)}

	ticker := time.NewTicker(tickerDuration)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fmt.Printf("%f/s\n", float64(*count.n)/(count.duration.Seconds()))
		}
	}()

	now := time.Now()
	for range ticker.C {
		count.n = Ptr(*count.n + 1)
		count.duration = Ptr(time.Since(now))
	}
}

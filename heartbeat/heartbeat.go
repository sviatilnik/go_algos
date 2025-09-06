package heartbeat

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Heartbeat(ctx context.Context) (<-chan int, <-chan interface{}) {
	result := make(chan int)
	beats := make(chan interface{})

	go func(ctx context.Context) {
		defer close(result)
		defer close(beats)

		pulse := time.NewTicker(time.Second)
		workTicker := time.NewTicker(3 * time.Second)

		defer pulse.Stop()
		defer workTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				beats <- time.Now().Unix()
			case <-workTicker.C:
				result <- rand.Int()
			}
		}

	}(ctx)

	return result, beats
}

func Sample() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	r, p := Heartbeat(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case res := <-r:
			fmt.Println(fmt.Sprintf("Got some result %d", res))
		case <-p:
			fmt.Println("Beat")
		}
	}
}

package throttle

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrTooManyRequests = errors.New("too many requests")

type Effector func(context.Context) (string, error)

func Throttle(e Effector, max uint, refill uint, d time.Duration) Effector {
	var tokens = max
	var once sync.Once

	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		// Пополнение токенов
		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}

						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return "", ErrTooManyRequests
		}

		tokens--

		return e(ctx)
	}
}

func Sample() {
	fn := func(ctx context.Context) (string, error) {
		return "Hello World!", nil
	}

	fn2 := func(ctx context.Context) (string, error) {
		return "Hello World2!", nil
	}

	throttle := Throttle(fn, 3, 1, 1*time.Nanosecond)
	throttle2 := Throttle(fn2, 1, 1, 500*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; i < 5; i++ {
		go func(i int) {
			r, e := throttle(ctx)
			fmt.Println(fmt.Sprintf("%d: %s %s %v", i, "throttle", r, e))
		}(i)

		go func(i int) {
			r, e := throttle2(ctx)
			fmt.Println(fmt.Sprintf("%d: %s %s %v", i, "throttle2", r, e))
		}(i)
	}

	time.Sleep(time.Second * 5)
}

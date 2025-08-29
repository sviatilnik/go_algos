package throttle

import (
	"context"
	"errors"
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

		if tokens < 0 {
			return "", ErrTooManyRequests
		}

		tokens--

		return e(ctx)
	}
}

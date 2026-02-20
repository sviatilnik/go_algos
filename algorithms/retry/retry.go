package retry

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Effector func(ctx context.Context) (string, error)

func Retry(fn Effector, attempts int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			res, err := fn(ctx)
			if err == nil || r >= attempts {
				return res, err
			}

			fmt.Printf("retry attempt #%d failed, retrying in %s\n", r, delay)

			select {
			case <-ctx.Done():
				return "", ctx.Err()
			// ждем delay
			case <-time.After(delay):
			}
		}
	}
}

func Sample() {
	var i = 0
	fn := func(ctx context.Context) (string, error) {
		i++
		if i == 1 {
			return "", errors.New("omg got some error. need to retry")
		}

		return strconv.Itoa(i), nil
	}

	retry := Retry(fn, 3, 500*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println(retry(ctx))
}

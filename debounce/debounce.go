package debounce

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

func Debounce(circuit Circuit, d time.Duration) Circuit {
	var m sync.Mutex
	var threshold time.Time // время исполнения следующего запроса
	var result string
	var err error

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			return result, err
		}

		result, err = circuit(ctx)

		return result, err
	}
}

func Sample() {
	circuit1 := func(ctx context.Context) (string, error) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			time.Sleep(1 * time.Second)
			return fmt.Sprintf("result was generated at %d", time.Now().Unix()), nil
		}
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	debounce := Debounce(circuit1, 5*time.Second)

	for range 10 {
		r, err := debounce(ctx)
		fmt.Println(r, err)
		time.Sleep(500 * time.Millisecond)
	}
}

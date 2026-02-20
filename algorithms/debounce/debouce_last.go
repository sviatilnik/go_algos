package debounce

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var threshold = time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var m sync.Mutex
	var once sync.Once

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()

		threshold = time.Now().Add(d)

		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)

			go func() {
				// Сброс после завершения ожидания threshold
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					// Проверяем, можно ли получать результат
					case <-ticker.C:
						m.Lock()
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							m.Unlock()
							return
						}
						m.Unlock()
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()

		})

		return result, err
	}
}

func LastSample() {
	circuit := func(ctx context.Context) (string, error) {
		return fmt.Sprintf("result was generated at %d", time.Now().Unix()), nil
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	debounce := DebounceLast(circuit, 1*time.Second)

	// Тут результата не будет, потому что таймаут 500мс, а threshold 1 секунда, т.е debounce будет каждый раз сбрасываться
	for range 5 {
		fmt.Println(debounce(ctx))
		time.Sleep(500 * time.Millisecond)
	}

	// Тут будет результат, потому что мы ждаи больше, чем threshold
	time.Sleep(2 * time.Second)
	fmt.Println(debounce(ctx))
}

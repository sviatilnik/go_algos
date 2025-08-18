package circuit_breaker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, errorsThreshold uint) Circuit {
	var failersCount = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		// Блокируем чтение, чтобы проверить можно ли обращаться к сервису
		m.RLock()
		// Вычисляем кол-во последних ошибок. Если они есть, то мы устанавливаем время следующей попытки
		delta := failersCount - int(errorsThreshold)
		if delta >= 0 {
			retryAt := lastAttempt.Add(time.Second * 5 << delta)
			if !time.Now().After(retryAt) {
				m.RUnlock()
				return "", errors.New("service unavailable")
			}
		}
		m.RUnlock()

		// Передаем запрос circuit
		resp, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			failersCount++
			return resp, err
		}

		failersCount = 0

		return resp, nil
	}
}

func Sample() {
	circuit1 := func(ctx context.Context) (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("OMG some error")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	breaker := Breaker(circuit1, 5)

	for range 10 {
		r, err := breaker(ctx)
		fmt.Println(r, err)
	}
}

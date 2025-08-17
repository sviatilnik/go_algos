package circuit_breaker

import (
	"context"
	"errors"
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
		if delta > 0 {
			retryAt := lastAttempt.Add(time.Second * 1 << delta)
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

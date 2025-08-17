package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"time"
)

func main() {

	circuit1 := func(ctx context.Context) (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("OMG some error")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	breaker := circuit_breaker.Breaker(circuit1, 5)

	for range 10 {
		r, err := breaker(ctx)
		fmt.Println(r, err)
	}
}

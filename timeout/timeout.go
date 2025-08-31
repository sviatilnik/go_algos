package timeout

import (
	"context"
	"fmt"
	"time"
)

type SlowFunction func(string) (string, error)

type WithContext func(context.Context, string) (string, error)

func Timeout(f SlowFunction) WithContext {
	return func(ctx context.Context, arg string) (string, error) {
		chres := make(chan string)
		cherr := make(chan error)

		go func() {
			res, err := f(arg)
			chres <- res
			cherr <- err
		}()

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case res := <-chres:
			return res, <-cherr
		}
	}
}

func Sample() {
	fn := func(str string) (string, error) {
		time.Sleep(1 * time.Second)

		return fmt.Sprintf("arg: %s", str), nil
	}

	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	timeout := Timeout(fn)
	res, err := timeout(ctxt, "some input")
	fmt.Println(res, err)
}

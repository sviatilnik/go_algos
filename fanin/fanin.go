package fanin

import (
	"fmt"
	"sync"
	"time"
)

func Funnel(sources ...<-chan int) <-chan int {
	dest := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(sources))
	for _, source := range sources {
		go func(c <-chan int) {
			defer wg.Done()

			for n := range c {
				dest <- n
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}

func Sample() {
	sources := make([]<-chan int, 0)

	for i := 0; i < 3; i++ {
		c := make(chan int)
		sources = append(sources, c)

		go func() {
			defer close(c)

			for j := 0; j < 5; j++ {
				c <- j
				time.Sleep(time.Second)
			}

		}()
	}

	dest := Funnel(sources...)
	for n := range dest {
		fmt.Println(n)
	}
}

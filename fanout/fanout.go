package fanout

import (
	"fmt"
	"sync"
)

func Split(source <-chan int, n int) []<-chan int {
	destination := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		c := make(chan int)
		destination = append(destination, c)

		go func() {
			defer close(c)

			for v := range source {
				c <- v
			}
		}()
	}

	return destination
}

func Sample() {
	source := make(chan int)
	destination := Split(source, 5)

	go func() {
		for i := 0; i <= 10; i++ {
			source <- i
		}

		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(destination))

	for i, v := range destination {
		go func(i int, dest <-chan int) {
			defer wg.Done()

			for val := range dest {
				fmt.Println(val)
			}
		}(i, v)
	}

	wg.Wait()
}

package generator

import "fmt"

func Generator(start, end uint) <-chan uint {
	result := make(chan uint)

	go func(r chan uint) {
		defer close(result)
		for i := start; i <= end; i++ {
			r <- i
		}
	}(result)

	return result
}

func Sample() {
	for r := range Generator(0, 10) {
		fmt.Println(r)
	}
}

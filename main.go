package main

import (
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"github.com/sviatilnik/go_algos/debounce"
)

func main() {
	// startCircuitBreakerSample()
	startDebounceSample()
}

func startDebounceSample() {
	debounce.Sample()
}

func startCircuitBreakerSample() {
	circuit_breaker.Sample()
}

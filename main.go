package main

import (
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"github.com/sviatilnik/go_algos/debounce"
)

func main() {
	// startCircuitBreakerSample()
	startDebounceFirstSample()
}

func startDebounceFirstSample() {
	debounce.FirstSample()
}

func startCircuitBreakerSample() {
	circuit_breaker.Sample()
}

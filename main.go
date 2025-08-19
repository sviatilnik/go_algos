package main

import (
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"github.com/sviatilnik/go_algos/debounce"
)

func main() {
	// startCircuitBreakerSample()
	// startDebounceFirstSample()
	startDebounceLastSample()
}

func startDebounceFirstSample() {
	debounce.FirstSample()
}

func startDebounceLastSample() {
	debounce.LastSample()
}

func startCircuitBreakerSample() {
	circuit_breaker.Sample()
}

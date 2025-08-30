package main

import (
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"github.com/sviatilnik/go_algos/debounce"
	"github.com/sviatilnik/go_algos/retry"
	"github.com/sviatilnik/go_algos/throttle"
)

func main() {
	// startCircuitBreakerSample()
	// startDebounceFirstSample()
	// startDebounceLastSample()
	// retrySample()
	throttleSample()
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

func retrySample() {
	retry.Sample()
}

func throttleSample() {
	throttle.Sample()
}

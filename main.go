package main

import (
	"github.com/sviatilnik/go_algos/circuit_breaker"
	"github.com/sviatilnik/go_algos/debounce"
	"github.com/sviatilnik/go_algos/fanin"
	"github.com/sviatilnik/go_algos/fanout"
	"github.com/sviatilnik/go_algos/generator"
	"github.com/sviatilnik/go_algos/heartbeat"
	"github.com/sviatilnik/go_algos/retry"
	"github.com/sviatilnik/go_algos/throttle"
	"github.com/sviatilnik/go_algos/timeout"
)

func main() {
	// startCircuitBreakerSample()
	// startDebounceFirstSample()
	// startDebounceLastSample()
	// retrySample()
	// throttleSample()
	// timeout.Sample()
	// FanInSample()
	// GeneratorSample()
	// HeartbeatSample()
	FanOutSample()
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

func TimeoutSample() {
	timeout.Sample()
}

func FanInSample() {
	fanin.Sample()
}

func GeneratorSample() {
	generator.Sample()
}

func HeartbeatSample() {
	heartbeat.Sample()
}

func FanOutSample() {
	fanout.Sample()
}

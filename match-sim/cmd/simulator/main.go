package main

import (
	"match-sim/mock"
)

func main() {
	// Load match from configs
	match := mock.Match1

	match.SimulateMatch(12345, false)
}

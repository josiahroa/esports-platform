package main

import (
	"match-sim/configs"
)

func main() {
	// Load match from configs
	match := configs.Match1

	match.SimulateMatch(13414123, false)
}

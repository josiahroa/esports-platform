package match

type Rules struct {
	RoundDurationSeconds             int
	RoundDurationSpikePlantedSeconds int
	RoundsPerHalf                    int
	Halves                           int
	SpikePlantTimeSeconds            int
	SpikeDefuseTimeSeconds           int
	RoundsToWin                      int
}

func DefaultRules() Rules {
	return Rules{
		RoundDurationSeconds:             100,
		RoundDurationSpikePlantedSeconds: 45,
		RoundsPerHalf:                    12,
		Halves:                           2,
		SpikePlantTimeSeconds:            4,
		SpikeDefuseTimeSeconds:           7,
		RoundsToWin:                      13,
	}
}

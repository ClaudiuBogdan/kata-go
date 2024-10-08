package twocrystalballs

import "math"

// TwoCrystalBalls finds the exact floor from which the crystal balls will break
func TwoCrystalBalls(breaks []bool) int {
	if len(breaks) == 0 {
		return -1
	}

	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	// We start from the first section
	i := jumpAmount

	for ; i < len(breaks) && !breaks[i]; i += jumpAmount {
	}

	i -= jumpAmount

	for ; i < len(breaks); i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}

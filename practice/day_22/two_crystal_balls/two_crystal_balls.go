package twocrystalballs

import "math"

// TwoCrystalBalls finds the exact floor from which the crystal balls will break
func TwoCrystalBalls(breaks []bool) int {
	incrementAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := incrementAmount

	for ; i < len(breaks) && !breaks[i]; i += incrementAmount {
	}

	i -= incrementAmount

	for ; i < len(breaks) && !breaks[i]; i++ {
	}

	if i == len(breaks) {
		return -1
	}

	return i
}

package twocrystalballs

import "math"

// TwoCrystalBalls finds the exact floor from which the crystal balls will break
func TwoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	// First ball: jump sqrt(n) floors at a time
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	// Second ball: linear search in the last sqrt(n) block
	i -= jumpAmount
	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	// If no breaking point found
	return -1
}

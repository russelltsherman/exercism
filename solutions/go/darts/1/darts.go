package darts

import (
    "math"
)

func Score(x, y float64) int {
	// Calculate radius (distance from center 0,0) using Pythagorean theorem
	radius := math.Hypot(x, y)

	switch {
	case radius <= 1:
		return 10 // Inner circle
	case radius <= 5:
		return 5  // Middle circle
	case radius <= 10:
		return 1  // Outer circle
	default:
		return 0  // Outside the target
	}
}

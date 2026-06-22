//go:build wasm

package canvas2d

import "math"

func compute_angle_difference(a float64, b float64) float64 {

	diff := a - b

	for diff > math.Pi {
		diff -= 2 * math.Pi
	}

	for diff < -math.Pi {
		diff += 2 * math.Pi
	}

	return diff

}


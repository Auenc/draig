package utils

import "math"

const GridSize = 12.0

func ProportionalDimension(ratio float64, dim int) int {
	if ratio == 0 {
		return 0
	}
	return int(math.Round((float64(dim) / GridSize) * ratio))
}

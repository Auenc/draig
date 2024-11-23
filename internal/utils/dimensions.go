package utils

const GridSize = 36

func ProportionalDimension(ratio int, dim int) int {
	if ratio == 0 {
		return 0
	}
	return (dim / GridSize) * ratio
}

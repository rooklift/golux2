package kit

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Make2dIntSlice(width int, height int, content int) [][]int {
	ret := make([][]int, width)
	for x := 0; x < width; x++ {
		ret[x] = make([]int, height)
		for y := 0; y < height; y++ {
			ret[x][y] = content
		}
	}
	return ret
}

func Make2dBoolSlice(width int, height int, content bool) [][]bool {
	ret := make([][]bool, width)
	for x := 0; x < width; x++ {
		ret[x] = make([]bool, height)
		for y := 0; y < height; y++ {
			ret[x][y] = content
		}
	}
	return ret
}

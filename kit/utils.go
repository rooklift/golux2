package kit

func bool_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func make_2d_int_slice(width int, height int, content int) [][]int {
	ret := make([][]int, width)
	for x := 0; x < width; x++ {
		ret[x] = make([]int, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			ret[x][y] = content
		}
	}
	return ret
}

func make_2d_bool_slice(width int, height int, content bool) [][]bool {
	ret := make([][]bool, width)
	for x := 0; x < width; x++ {
		ret[x] = make([]bool, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			ret[x][y] = content
		}
	}
	return ret
}

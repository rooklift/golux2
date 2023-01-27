package kit

import "strings"

func bool_to_int(b bool) int { if b { return 1 }
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

func BoardASCII() string {
	var elements []string
	board := GetBoard()
	for y := 0; y < Height(); y++ {
		for x := 0; x < Width(); x++ {
			s := "  "
			if board.FactoryOccupancy[x][y] > -1 {
				s = " @"
			} else if board.Ore[x][y] > 0 {
				s = " O"
			} else if board.Ice[x][y] > 0 {
				s = " X"
			} else if board.Rubble[x][y] > 0 {
				s = " ."
			}
			elements = append(elements, s)
		}
		elements = append(elements, "\n")
	}
	return strings.Join(elements, "")
}

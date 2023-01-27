package kit

import "strings"

func bool_to_int(b bool) int { if b { return 1 }
	return 0
}

func BoardASCII() string {
	var elements []string
	board := GetBoard()
	for y := 0; y < Height(); y++ {
		for x := 0; x < Width(); x++ {
			s := "  "
			if board.Ore[x][y] > 0 {
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

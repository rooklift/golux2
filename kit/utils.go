package kit

import "strings"

func bool_to_int(b bool) int { if b { return 1 }
	return 0
}

func WorldASCII() string {

	var elements []string

	for y := 0; y < Height(); y++ {
		for x := 0; x < Width(); x++ {
			elements = append(elements, " .")
		}
		elements = append(elements, "\n")
	}

	return strings.Join(elements, "")

}


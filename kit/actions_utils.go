package kit

func bool_to_int(b bool) int {
	if b { return 1 }
	return 0
}

func MoveAction(direction int, replace bool, iterations int) []int {
	return []int{0, direction, 0, 0, bool_to_int(replace), iterations}
}

func ActionQueue(args ...[]int) [][]int {
	var ret [][]int
	for _, item := range args {
		ret = append(ret, item)
	}
	return ret
}

package kit

// A robot action is a length-6 array as follows:
// [0] type
// [1] direction
// [2] resource
// [3] amount
// [4] send_to_back
// [5] iterations

func bool_to_int(b bool) int {
	if b { return 1 }
	return 0
}

func ActionQueue(args ...[6]int) [][6]int {
	var ret [][6]int
	for _, item := range args {
		ret = append(ret, item)
	}
	return ret
}

func Action(atype int, direction int, resource int, amount int, send_to_back bool, iterations int) [6]int {
	return [6]int{atype, direction, resource, amount, bool_to_int(send_to_back), iterations}
}

/*
func Move(direction int, send_to_back bool, iterations int) [6]int {
	return [6]int{MOVE, direction, 0, 0, bool_to_int(send_to_back), iterations}
}
func Transfer(direction int, resource int, amount int, send_to_back bool, iterations int) [6]int {
	return [6]int{TRANSFER, direction, resource, amount, bool_to_int(send_to_back), iterations}
}
func Pickup(resource int, amount int, send_to_back bool, iterations int) [6]int {
	return [6]int{PICKUP, 0, resource, amount, bool_to_int(send_to_back), iterations}
}
func Dig(send_to_back bool, iterations int) [6]int {
	return [6]int{DIG, 0, 0, 0, bool_to_int(send_to_back), iterations}
}
func SelfDestruct() [6]int {
	return [6]int{SELFDESTRUCT, 0, 0, 0, 0, 0}
}
func Recharge(amount int, send_to_back bool, iterations int) [6]int {
	return [6]int{RECHARGE, 0, 0, amount, bool_to_int(send_to_back), iterations}
}
*/
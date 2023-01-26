package kit

import "fmt"

func Bid(faction string, bid int) {
	bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func PlaceFactory(x int, y int, metal int, water int) {
	placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

// ------------------------------------------------------------------------------------------------

func FactoryAct(uid string, action int) {
	factory_actions[uid] = action
}

func FactoryCancel(uid string) {
	delete(factory_actions, uid)
}

func (self *Factory) Act(action int) {					// Method is just a convenient shorthand for the above.
	FactoryAct(self.UnitId, action)
}

func (self *Factory) Cancel() {
	FactoryCancel(self.UnitId)
}

// ------------------------------------------------------------------------------------------------
// A robot action is a length-6 array as follows:
// [0] type
// [1] direction
// [2] resource
// [3] amount
// [4] send_to_back
// [5] iterations

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

func RobotSetQueue(uid string, args ...[6]int) {
	var queue [][6]int
	for _, item := range args {
		queue = append(queue, item)
	}
	robot_actions[uid] = queue
}

func RobotCancel(uid string) {
	delete(robot_actions, uid)
}

func (self *Unit) SetQueue(args ...[6]int) {			// Method is just a convenient shorthand for the above.
	RobotSetQueue(self.UnitId, args...)
}

func (self *Unit) Cancel() {
	RobotCancel(self.UnitId)
}

// ------------------------------------------------------------------------------------------------

func bool_to_int(b bool) int { if b { return 1 }
	return 0
}

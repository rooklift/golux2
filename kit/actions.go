package kit

import "fmt"

func (self *Frame) Bid(faction string, bid int) {
	self.bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func (self *Frame) PlaceFactory(x int, y int, metal int, water int) {
	self.placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

// ------------------------------------------------------------------------------------------------

func (self *Factory) Act(action int) {
	self.Frame.factory_actions[self.UnitId] = action
}

func (self *Factory) Cancel() {
	delete(self.Frame.factory_actions, self.UnitId)
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

func (self *Unit) SetQueue(args ...[6]int) {
	var queue [][6]int
	for _, item := range args {
		queue = append(queue, item)
	}
	self.Frame.unit_actions[self.UnitId] = queue
}

func (self *Unit) Cancel() {
	delete(self.Frame.unit_actions, self.UnitId)
}

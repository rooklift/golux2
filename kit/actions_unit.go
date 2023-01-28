package kit

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

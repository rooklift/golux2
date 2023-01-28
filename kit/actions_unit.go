package kit

func Action(atype int, direction int, resource int, amount int, recycle int, iterations int) [6]int {
	return [6]int{atype, direction, resource, amount, recycle, iterations}
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

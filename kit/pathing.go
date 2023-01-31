package kit

func (self *Unit) NaiveTrip(other XYer) []Action {

	x1, y1 := self.XY()
	x2, y2 := other.XY()

	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 && dy == 0 {
		return nil
	}

	var ret []Action

	if dx > 0 {
		ret = append(ret, Action{MOVE, RIGHT, 0, 0, 0, abs(dx)})
	} else if dx < 0 {
		ret = append(ret, Action{MOVE, LEFT, 0, 0, 0, abs(dx)})
	}

	if dy > 0 {
		ret = append(ret, Action{MOVE, DOWN, 0, 0, 0, abs(dy)})
	} else if dy < 0 {
		ret = append(ret, Action{MOVE, UP, 0, 0, 0, abs(dy)})
	}

	return ret
}

package kit

func (self *Frame) IceAt(xy XYer) bool {
	x, y := xy.XY()
	if x < 0 || x >= self.Width() || y < 0 || y >= self.Height() {
		return false
	}
	return self.GetBoard().Ice[x][y] > 0
}

func (self *Frame) OreAt(xy XYer) bool {
	x, y := xy.XY()
	if x < 0 || x >= self.Width() || y < 0 || y >= self.Height() {
		return false
	}
	return self.GetBoard().Ore[x][y] > 0
}

func (self *Frame) RubbleAt(xy XYer) int {
	x, y := xy.XY()
	if x < 0 || x >= self.Width() || y < 0 || y >= self.Height() {
		return 0
	}
	return self.GetBoard().Rubble[x][y]
}

func (self *Frame) AllIce() []Pos {
	var ret []Pos
	for x := 0; x < self.Width(); x++ {
		for y := 0; y < self.Height(); y++ {
			if self.IceAt(Pos{x, y}) {
				ret = append(ret, Pos{x, y})
			}
		}
	}
	return ret
}

func (self *Frame) AllOre() []Pos {
	var ret []Pos
	for x := 0; x < self.Width(); x++ {
		for y := 0; y < self.Height(); y++ {
			if self.OreAt(Pos{x, y}) {
				ret = append(ret, Pos{x, y})
			}
		}
	}
	return ret
}

func (self *Frame) AllRubble() []Pos {
	var ret []Pos
	for x := 0; x < self.Width(); x++ {
		for y := 0; y < self.Height(); y++ {
			n := self.RubbleAt(Pos{x, y})
			if n > 0 {
				ret = append(ret, Pos{x, y})
			}
		}
	}
	return ret
}

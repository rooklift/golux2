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
	if self.cache.all_ice == nil {
		self.cache.all_ice = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.IceAt(Pos{x, y}) {
					self.cache.all_ice = append(self.cache.all_ice, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.cache.all_ice))
	copy(ret, self.cache.all_ice)
	return ret
}

func (self *Frame) AllOre() []Pos {
	if self.cache.all_ore == nil {
		self.cache.all_ore = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.OreAt(Pos{x, y}) {
					self.cache.all_ore = append(self.cache.all_ore, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.cache.all_ore))
	copy(ret, self.cache.all_ore)
	return ret
}

func (self *Frame) AllRubble() []Pos {
	if self.cache.all_rubble == nil {
		self.cache.all_rubble = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.RubbleAt(Pos{x, y}) > 0 {
					self.cache.all_rubble = append(self.cache.all_rubble, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.cache.all_rubble))
	copy(ret, self.cache.all_rubble)
	return ret
}

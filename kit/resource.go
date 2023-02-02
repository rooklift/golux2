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
	if self.all_ice_cache == nil {
		self.all_ice_cache = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.IceAt(Pos{x, y}) {
					self.all_ice_cache = append(self.all_ice_cache, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.all_ice_cache))
	copy(ret, self.all_ice_cache)
	return ret
}

func (self *Frame) AllOre() []Pos {
	if self.all_ore_cache == nil {
		self.all_ore_cache = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.OreAt(Pos{x, y}) {
					self.all_ore_cache = append(self.all_ore_cache, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.all_ore_cache))
	copy(ret, self.all_ore_cache)
	return ret
}

func (self *Frame) AllRubble() []Pos {
	if self.all_rubble_cache == nil {
		self.all_rubble_cache = []Pos{}			// Ensure not nil even if it ends up being empty
		for x := 0; x < self.Width(); x++ {
			for y := 0; y < self.Height(); y++ {
				if self.RubbleAt(Pos{x, y}) > 0 {
					self.all_rubble_cache = append(self.all_rubble_cache, Pos{x, y})
				}
			}
		}
	}
	ret := make([]Pos, len(self.all_rubble_cache))
	copy(ret, self.all_rubble_cache)
	return ret
}

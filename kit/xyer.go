package kit

type XYer interface {
	X() int
	Y() int
	XY() (int, int)
	GetPos() Pos
}

// We generally use Pos in a pass-by-value way, but pass units and factories by pointer...

func (pos Pos) X() int { return pos[0] }
func (pos Pos) Y() int { return pos[1] }
func (pos Pos) XY() (int, int) { return pos[0], pos[1] }
func (pos Pos) GetPos() Pos { return pos }

func (self *Unit) X() int { return self.Pos[0] }
func (self *Unit) Y() int { return self.Pos[1] }
func (self *Unit) XY() (int, int) { return self.Pos[0], self.Pos[1] }
func (self *Unit) GetPos() Pos { return self.Pos }

func (self *Factory) X() int { return self.Pos[0] }
func (self *Factory) Y() int { return self.Pos[1] }
func (self *Factory) XY() (int, int) { return self.Pos[0], self.Pos[1] }
func (self *Factory) GetPos() Pos { return self.Pos }

// The only reason any of this exists...

func Dist(a XYer, b XYer) int {
	x1, y1 := a.XY()
	x2, y2 := b.XY()
	return abs(x2 - x1) + abs(y2 - y1)
}

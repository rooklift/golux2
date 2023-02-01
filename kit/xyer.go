package kit

type XYer interface {
	XY() (int, int)
}

// We generally use Pos in a pass-by-value way, but pass units and factories by pointer...

func (pos Pos) XY() (int, int) { return pos.X, pos.Y }
func (self *Unit) XY() (int, int) { return self.Pos.X, self.Pos.Y }
func (self *Factory) XY() (int, int) { return self.Pos.X, self.Pos.Y }

// The only reason any of this exists...

func Dist(a XYer, b XYer) int {
	x1, y1 := a.XY()
	x2, y2 := b.XY()
	return abs(x2 - x1) + abs(y2 - y1)
}

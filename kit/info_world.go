package kit

import "strings"

func (self *Frame) GetBoard() *Board {
	return self.Obs.Board
}

func (self *Frame) RealStep() int {
	return self.Obs.RealEnvSteps
}

func (self *Frame) Width() int {
	return len(self.GetBoard().Rubble)
}

func (self *Frame) Height() int {
	return len(self.GetBoard().Rubble[0])
}

func (self *Frame) BoardASCII() string {				// For logging
	var elements []string
	board := self.GetBoard()
	for y := 0; y < self.Height(); y++ {
		for x := 0; x < self.Width(); x++ {
			s := "  "
			if board.FactoryOccupancy[x][y] > -1 {
				s = " @"
			} else if board.Ore[x][y] > 0 {
				s = " O"
			} else if board.Ice[x][y] > 0 {
				s = " X"
			} else if board.Rubble[x][y] > 0 {
				s = " ."
			}
			elements = append(elements, s)
		}
		elements = append(elements, "\n")
	}
	return strings.Join(elements, "")
}

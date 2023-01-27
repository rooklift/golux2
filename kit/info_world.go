package kit

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

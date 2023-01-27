package kit

func GetMsg() *Message {
	return msg
}

func GetBoard() *Board {
	return msg.Obs.Board
}

func RealStep() int {
	return msg.Obs.RealEnvSteps
}

func Width() int {
	return len(GetBoard().Rubble)
}

func Height() int {
	return len(GetBoard().Rubble[0])
}

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

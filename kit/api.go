package kit

func MyPlayerId() string {
	return msg.Player
}

func MyTeam() *Team {
	return msg.Obs.Teams[MyPlayerId()]
}

func GetMsg() *Message {
	return msg
}

func GetBoard() *Board {
	return msg.Obs.Board
}

func CanPlaceFactory() bool {
	team := MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && msg.Step % 2 == 1) || (!team.PlaceFirst && msg.Step % 2 == 0)
}

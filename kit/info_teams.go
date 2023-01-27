package kit

func MyPlayerId() string {
	ret := msg.Player
	if ret != "player_0" && ret != "player_1" {
		panic("player was neither player_0 nor player_1 - this violates some assumptions")
	}
	return ret
}

func TheirPlayerId() string {
	if MyPlayerId() == "player_0" {
		return "player_1"
	}
	return "player_0"
}

func MyTeam() *Team {
	return msg.Obs.Teams[MyPlayerId()]
}

func TheirTeam() *Team {
	return msg.Obs.Teams[TheirPlayerId()]
}

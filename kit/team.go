package kit

func (self *Frame) MyPlayerId() string {
	ret := self.Player
	if ret != "player_0" && ret != "player_1" {
		panic("player was neither player_0 nor player_1 - this violates some assumptions")
	}
	return ret
}

func (self *Frame) TheirPlayerId() string {
	if self.MyPlayerId() == "player_0" {
		return "player_1"
	}
	return "player_0"
}

func (self *Frame) MyTeam() *Team {
	return self.Obs.Teams[self.MyPlayerId()]
}

func (self *Frame) TheirTeam() *Team {
	return self.Obs.Teams[self.TheirPlayerId()]
}

func (self *Frame) MyPlayerInt() int {
	return self.MyTeam().TeamId
}

func (self *Frame) TheirPlayerInt() int {
	return self.TheirTeam().TeamId
}

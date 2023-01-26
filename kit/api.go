package kit

import "fmt"

var bid_string string
var placement_string string

func Bid(faction string, bid int) {
	bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func PlaceFactory(x int, y int, metal int, water int) {
	placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

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

package kit

import "sort"

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

func RealStep() int {
	return msg.Obs.RealEnvSteps
}

func CanPlaceFactory() bool {
	team := MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && msg.Step % 2 == 1) || (!team.PlaceFirst && msg.Step % 2 == 0)
}

func MyFactories() []*Factory {						// We sort keys for deterministic order
	my_fact_map := msg.Obs.Factories[MyPlayerId()]
	var keys []string
	for key, _ := range my_fact_map {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var ret []*Factory
	for _, key := range keys {
		ret = append(ret, my_fact_map[key])
	}
	return ret
}

func MyUnits() []*Unit {							// We sort keys for deterministic order
	my_unit_map := msg.Obs.Units[MyPlayerId()]
	var keys []string
	for key, _ := range my_unit_map {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var ret []*Unit
	for _, key := range keys {
		ret = append(ret, my_unit_map[key])
	}
	return ret
}

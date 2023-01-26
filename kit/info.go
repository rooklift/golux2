package kit

import "sort"

func MyPlayerId() string {
	return msg.Player
}

func TheirPlayerId() string {
	_, ok := msg.Obs.Teams[MyPlayerId()]		// Sanity check that msg.Obs.Teams at least contains my id
	if (!ok) {
		panic("TheirPlayerId(): sanity check failed")
	}
	for key, _ := range msg.Obs.Teams {
		if key != MyPlayerId() {
			return key
		}
	}
	panic("TheirPlayerId(): failed to get result")
}

func MyTeam() *Team {
	return msg.Obs.Teams[MyPlayerId()]
}

func TheirTeam() *Team {
	return msg.Obs.Teams[TheirPlayerId()]
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

func get_factories(playerid string) []*Factory {					// We sort keys for deterministic order
	fact_map := msg.Obs.Factories[playerid]
	var keys []string
	for key, _ := range fact_map {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var ret []*Factory
	for _, key := range keys {
		ret = append(ret, fact_map[key])
	}
	return ret
}

func MyFactories() []*Factory {	
	return get_factories(MyPlayerId())
}

func TheirFactories() []*Factory {
	return get_factories(TheirPlayerId())
}

func get_units(playerid string) []*Unit {							// We sort keys for deterministic order
	unit_map := msg.Obs.Units[playerid]
	var keys []string
	for key, _ := range unit_map {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var ret []*Unit
	for _, key := range keys {
		ret = append(ret, unit_map[key])
	}
	return ret
}

func MyUnits() []*Unit {
	return get_units(MyPlayerId())
}

func TheirUnits() []*Unit {
	return get_units(TheirPlayerId())
}


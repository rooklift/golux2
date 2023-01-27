package kit

import "sort"

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

func AllFactories() []*Factory {
	return append(MyFactories(), TheirFactories()...)				// Safe since these slices are constructed in the moment
}

func CanPlaceFactory() bool {
	team := MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && msg.Step % 2 == 1) || (!team.PlaceFirst && msg.Step % 2 == 0)
}
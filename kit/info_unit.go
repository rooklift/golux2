package kit

import "sort"

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

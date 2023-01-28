package kit

import "sort"

func (self *Frame) get_units(playerid string) []*Unit {			// We sort keys for deterministic order
	unit_map := self.Obs.Units[playerid]
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

func (self *Frame) MyUnits() []*Unit {
	return self.get_units(self.MyPlayerId())
}

func (self *Frame) TheirUnits() []*Unit {
	return self.get_units(self.TheirPlayerId())
}

func (self *Frame) AllUnits() []*Unit {
	return append(self.MyUnits(), self.TheirUnits()...)			// Safe since these slices are constructed in the moment
}

// ------------------------------------------------------------------------------------------------

func (self *Unit) X() int {
	return self.Pos[0]
}

func (self *Unit) Y() int {
	return self.Pos[1]
}

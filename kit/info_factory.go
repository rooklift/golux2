package kit

import "sort"

func (self *Frame) get_factories(playerid string) []*Factory {		// We sort keys for deterministic order
	fact_map := self.Obs.Factories[playerid]
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

func (self *Frame) MyFactories() []*Factory {	
	return self.get_factories(self.MyPlayerId())
}

func (self *Frame) TheirFactories() []*Factory {
	return self.get_factories(self.TheirPlayerId())
}

func (self *Frame) AllFactories() []*Factory {
	return append(self.MyFactories(), self.TheirFactories()...)		// Safe since these slices are constructed in the moment
}

func (self *Frame) CanPlaceFactory() bool {
	team := self.MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && self.Step % 2 == 1) || (!team.PlaceFirst && self.Step % 2 == 0)
}

// ------------------------------------------------------------------------------------------------

func (self *Factory) X() int {
	return self.Pos[0]
}

func (self *Factory) Y() int {
	return self.Pos[1]
}

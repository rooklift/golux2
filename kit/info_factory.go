package kit

import (
	"math/rand"
	"sort"
)

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

// ------------------------------------------------------------------------------------------------

func (self *Frame) CanPlaceFactory() bool {
	team := self.MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && self.Step % 2 == 1) || (!team.PlaceFirst && self.Step % 2 == 0)
}

func (self *Frame) PotentialSpawns() [][2]int {
	var ret [][2]int
	board := self.GetBoard()
	for x := 0; x < self.Width(); x++ {
		for y := 0; y < self.Height(); y++ {
			if board.ValidSpawnsMask[x][y] {
				ret = append(ret, [2]int{x, y})
			}
		}
	}
	return ret
}

func (self *Frame) RandomSpawn() [2]int {
	pot := self.PotentialSpawns()
	return pot[rand.Intn(len(pot))]
}

// ------------------------------------------------------------------------------------------------

func (self *Factory) X() int {
	return self.Pos[0]
}

func (self *Factory) Y() int {
	return self.Pos[1]
}

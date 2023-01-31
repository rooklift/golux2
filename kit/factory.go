package kit

import (
	"fmt"
	"math/rand"
	"sort"
)

func (self *Factory) Act(action int) {
	self.Request = action
}

func (self *Factory) ClearRequest() {
	self.Request = -1
}

func (self *Factory) HasRequest() bool {
	return self.Request >= 0
}

func (self *Factory) IsMine() bool {
	return self.TeamId == self.Frame.MyPlayerInt()
}

// ------------------------------------------------------------------------------------------------

func (self *Frame) Bid(faction string, bid int) {
	self.bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}", faction, bid)
}

func (self *Frame) PlaceFactory(pos Pos, metal int, water int) {
	self.placement_string = fmt.Sprintf("{\"spawn\": [%d,%d], \"metal\": %d, \"water\": %d}", pos.X(), pos.Y(), metal, water)
}

// ------------------------------------------------------------------------------------------------

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

func (self *Frame) FactoryByStrain(n int) *Factory {
	for _, factory := range self.AllFactories() {
		if factory.StrainId == n {
			return factory
		}
	}
	return nil
}

func (self *Frame) FactoryAt(xy XYer) *Factory {
	x, y := xy.XY()
	strain := self.Obs.Board.FactoryOccupancy[x][y]
	if strain == -1 {
		return nil
	}
	ret := self.FactoryByStrain(strain)
	return ret
}

func (self *Frame) CanPlaceFactory() bool {
	team := self.MyTeam()
	if team.FactoriesToPlace == 0 {
		return false
	}
	return (team.PlaceFirst && self.Step % 2 == 1) || (!team.PlaceFirst && self.Step % 2 == 0)
}

func (self *Frame) PotentialSpawns() []Pos {
	var ret []Pos
	board := self.GetBoard()
	for x := 0; x < self.Width(); x++ {
		for y := 0; y < self.Height(); y++ {
			if board.ValidSpawnsMask[x][y] {
				ret = append(ret, Pos{x, y})
			}
		}
	}
	return ret
}

func (self *Frame) RandomSpawn() Pos {
	pots := self.PotentialSpawns()
	return pots[rand.Intn(len(pots))]
}

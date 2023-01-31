package kit

import "sort"

func (self *Unit) AddToRequest(args ...Action) {

	// Exceeding 20 length is checked at the time of sending, no need to do it here.

	if len(args) == 0 {
		return
	}
	for _, action := range args {
		if action[4] < 0 {
			Log("%v - attempted to submit action with recycle == %d", self.UnitId, action[4])
			action[4] = 0
		}
		if action[5] < 1 {
			Log("%v - attempted to submit action with n == %d", self.UnitId, action[5])
			action[5] = 1
		}
		self.Request = append(self.Request, action)
	}
}

func (self *Unit) ClearRequest() {
	self.Request = nil
}

func (self *Unit) HasRequest() bool {
	return len(self.Request) > 0
}

func (self *Unit) IsMine() bool {
	return self.TeamId == self.Frame.MyPlayerInt()
}

// ------------------------------------------------------------------------------------------------

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

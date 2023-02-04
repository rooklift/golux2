package kit

import "sort"

func (self *Unit) BuildRequest(args ...Action) {

	if self.Request == nil {					// Calling this function (even with no args)
		self.Request = []Action{}				// makes the request not nil, by design.
	}

	// Exceeding 20 length is checked at the time of sending, no need to do it here.

	for _, action := range args {
		if action.Recycle < 0 {
			Log("%v - attempted to submit action with recycle == %d", self.UnitId, action.Recycle)
			action.Recycle = 0
		}
		if action.N < 1 {
			Log("%v - attempted to submit action with n == %d", self.UnitId, action.N)
			action.N = 1
		}
		self.Request = append(self.Request, action)
	}
}

func (self *Unit) ClearRequest() {
	self.Request = nil
}

func (self *Unit) HasRequest() bool {
	return self.Request != nil
}

func (self *Unit) PowerAfterRequest() int {						// This assumes there is a request.
	return self.Power - self.Frame.GetCfg().Robots[self.UnitType].ActionQueuePowerCost
}

func (self *Unit) CanAcceptRequest() bool {
	return self.PowerAfterRequest() >= 0
}

func (self *Unit) IsMine() bool {
	return self.TeamId == self.Frame.MyPlayerInt()
}

// ------------------------------------------------------------------------------------------------

func (self *Frame) get_units(playerid string) []*Unit {
	var ret []*Unit
	for _, unit := range self.Obs.Units[playerid] {
		ret = append(ret, unit)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].UnitInt < ret[j].UnitInt
	})
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

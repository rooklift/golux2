package main

import "golux2/kit"

// Register what 3 functions should be called, for the 3 phases of the game...

func main() {
	kit.Run(Bidder, Placer, AI)
}

// ------------------------------------------------------------------------------------------------

func Bidder(f *kit.Frame) {
	f.Bid("MotherMars", 0)
}

func Placer(f *kit.Frame) {
	if f.CanPlaceFactory() {
		f.PlaceFactory(f.RandomSpawn(), 150, 150)
	}
}

func AI(f *kit.Frame) {

	if f.RealStep() == 0 {
		for _, factory := range f.MyFactories() {
			factory.Act(kit.BUILD_LIGHT)						// Other actions are BUILD_HEAVY and WATER_LICHEN
		}
	}

	// Each turn, each unit can request to change their action queue (though this request can fail if they have no power).
	// The new action queue can have many actions. To make this easy to do, we provide a helper method, BuildRequest(),
	// which allows a new action queue to be built up. However, it can also accept more than 1 argument at a time.
	//
	// If necessary, you can call unit.ClearRequest() to delete whatever you have built up (this does not remove the
	// unit's existing action queue, however).

	if f.RealStep() == 1 {
		for _, unit := range f.MyUnits() {
			needed_power := f.GetCfg().Robots["LIGHT"].BatteryCapacity - unit.Power
			unit.BuildRequest(kit.PickupAction(kit.POWER, needed_power, 0, 1))
			unit.BuildRequest(kit.MoveAction(kit.LEFT, 2, 1))
			unit.BuildRequest(kit.MoveAction(kit.UP, 2, 1))
			// unit.ClearRequest()								// Try uncommenting this...
			unit.BuildRequest(kit.MoveAction(kit.RIGHT, 2, 2))
			unit.BuildRequest(kit.MoveAction(kit.DOWN, 2, 2))
		}
	}

	if f.RealStep() == 100 {
		for _, unit := range f.MyUnits() {
			naive_path := unit.NaiveTrip(kit.Pos{24,24})		// This is likely 2 actions long.
			unit.BuildRequest(naive_path...)
		}
	}
	
	// If you don't make any calls to BuildRequest() then nothing is sent for that unit, and so its old queue can survive.
}

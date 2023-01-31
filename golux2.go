package main

import "golux2/kit"

type Action = kit.Action
type Frame = kit.Frame
type Pos = kit.Pos

// Register what 3 functions should be called, for the 3 phases of the game...

func main() {
	kit.Run(Bidder, Placer, AI)
}

func Bidder(f *Frame) {
	f.Bid("MotherMars", 0)
}

func Placer(f *Frame) {
	if f.CanPlaceFactory() {
		f.PlaceFactory(f.RandomSpawn(), 150, 150)
	}
}

func AI(f *Frame) {

	if f.RealStep() == 0 {
		for _, factory := range f.MyFactories() {
			factory.Act(kit.BUILD_LIGHT)
		}
	}

	// Each turn, each unit can request to change their action queue (though this request can fail if they have no power).
	// The new action queue can have many actions. To make this easy to do, we provide a helper method, AddToRequest(),
	// which allows a new action queue to be built up. However, it can also accept more than 1 argument at a time.

	if f.RealStep() == 1 {
		for _, unit := range f.MyUnits() {
			unit.AddToRequest(Action{kit.PICKUP, kit.CENTRE, kit.POWER, f.GetCfg().Robots["LIGHT"].BatteryCapacity - unit.Power, 0, 1})
			unit.AddToRequest(Action{kit.MOVE, kit.LEFT, 0, 0, 2, 1})		// If you understand these numbers
			unit.AddToRequest(Action{kit.MOVE, kit.UP, 0, 0, 2, 1})			// you understand the action system...
			unit.AddToRequest(Action{kit.MOVE, kit.RIGHT, 0, 0, 2, 2})
			unit.AddToRequest(Action{kit.MOVE, kit.DOWN, 0, 0, 2, 2})
		}
	}

	// Note also that AddToRequest() is not modifying your existing ActionQueue (which can persist for many turns)
	// but replacing it with a completely new one - which, however, is built up with calls to AddToRequest().

	if f.RealStep() == 100 {
		for _, unit := range f.MyUnits() {
			unit.AddToRequest(unit.NaiveTrip(kit.Pos{24,24})...)			// The NaiveTrip() method returns a slice of actions
		}
	}
	
	// If you don't make any calls to AddToRequest() then nothing is sent for that unit, and so its old queue can survive.
}

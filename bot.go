package main

import (
	"golux2/kit"
)

// Tell the kit what 3 functions to use...

func main() {
	kit.Run(bidder, factory_placer, main_ai)
}

// Each turn, the correct function (of the 3) will be called
// and supplied with the current frame...

func bidder(f *kit.Frame) {
	kit.CreateLog(f.MyPlayerId() + ".log")
	kit.LogActions(true)
	f.Bid("MotherMars", 0)
}

func factory_placer(f *kit.Frame) {
	if f.CanPlaceFactory() {
		pos := f.RandomSpawn()
		f.PlaceFactory(pos[0], pos[1], 150, 150)
	}
}

func main_ai(f *kit.Frame) {
	if f.RealStep() == 0 {
		kit.Log(f.BoardASCII())
		for _, factory := range f.MyFactories() {
			factory.Act(kit.LIGHT)
		}
	}
	if f.RealStep() == 1 {
		for _, unit := range f.MyUnits() {
			unit.SetQueue(
				kit.Action(kit.MOVE, kit.LEFT, 0, 0, 2, 1),		// If you understand these numbers
				kit.Action(kit.MOVE, kit.UP, 0, 0, 2, 1),		// you understand the action system
				kit.Action(kit.MOVE, kit.RIGHT, 0, 0, 2, 2),	// though note they are only valid
				kit.Action(kit.MOVE, kit.DOWN, 0, 0, 2, 2),		// after Lux 2.1.0
			)
		}
	}
}

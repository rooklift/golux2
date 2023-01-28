package main

import (
	"golux2/kit"
)

// Tell the kit what 3 functions to use...

func main() {
	kit.Run(bidder, factory_placer, main_ai)
}

// Each of these can access state and setup actions...

func bidder(f *kit.Frame) {
	kit.CreateLog(f.MyPlayerId() + ".log")
	f.Bid("MotherMars", 0)
}

func factory_placer(f *kit.Frame) {
	if f.CanPlaceFactory() {
		board := f.GetBoard()
		for y := 0; y < f.Height(); y++ {
			for x := 0; x < f.Width(); x++ {
				if board.ValidSpawnsMask[x][y] {
					f.PlaceFactory(x, y, 150, 150)
					return
				}
			}
		}
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
				kit.Action(kit.MOVE, kit.RIGHT, 0, 0, 0, 1),
				kit.Action(kit.MOVE, kit.DOWN, 0, 0, 0, 1),
				kit.Action(kit.MOVE, kit.LEFT, 0, 0, 1, 2),
				kit.Action(kit.MOVE, kit.UP, 0, 0, 1, 2),
				kit.Action(kit.MOVE, kit.RIGHT, 0, 0, 1, 2),
				kit.Action(kit.MOVE, kit.DOWN, 0, 0, 1, 2),
			)
		}
	}
}

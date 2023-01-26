package main

import (
	"golux2/kit"
)

// Tell the kit what 3 functions to use...

func main() {
	kit.Run(bidder, factory_placer, main_ai)
}

// Each of these can access state and setup actions...

func bidder() {
	kit.Bid("MotherMars", 0)
}

func factory_placer() {
	if kit.CanPlaceFactory() {
		board := kit.GetBoard()
		for y := 0; y < 48; y++ {
			for x := 0; x < 48; x++ {
				if board.ValidSpawnsMask[x][y] {
					kit.PlaceFactory(x, y, 150, 150)
					return
				}
			}
		}
	}
}

func main_ai() {

	if kit.GetMsg().Obs.RealEnvSteps < 2 {
		for _, factory := range kit.MyFactories() {
			factory.Act(0)
		}
	}

	if kit.GetMsg().Obs.RealEnvSteps == 1 {
		for _, unit := range kit.MyUnits() {
			unit.Act([][]int{{0, 2, 0, 0, 0, 1}})
		}
	}
}

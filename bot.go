package main

import (
	"golux2/kit"
)

func main() {
	kit.Run(bidder, factory_placer, main_ai)
}

func bidder() {
	kit.Bid("AlphaStrike", 0)
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
	return
}

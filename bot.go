package main

import (
	"fmt"
	"golux2/kit"
)

func main() {
	kit.Register(bidder, factory_placer, main_ai)
}

func bidder() {
	kit.Bid("AlphaStrike", 0)
}

func factory_placer() {
	if kit.CanPlaceFactory() {
		board := kit.GetBoard()
		spawn_loop:
		for y := 0; y < 48; y++ {
			for x := 0; x < 48; x++ {
				if board.ValidSpawnsMask[x][y] {
					kit.PlaceFactory(x, y, 150, 150)
					break spawn_loop
				}
			}
		}
	}
}

func main_ai() {
	fmt.Printf("{}\n")
}

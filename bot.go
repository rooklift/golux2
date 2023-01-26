package main

import (
	"fmt"
	"golux2/kit"
)

func main() {
	kit.Register(bidder, factory_placer, main_ai)
}

func bidder() {
	fmt.Printf("{\"faction\": \"AlphaStrike\", \"bid\": 0}\n")
}

func factory_placer() {
	if kit.CanPlaceFactory() {
		msg := kit.GetMsg()
		spawn_loop:
		for y := 0; y < 48; y++ {
			for x := 0; x < 48; x++ {
				if msg.Obs.Board.ValidSpawnsMask[x][y] {
					fmt.Printf("{\"spawn\": [%d, %d], \"metal\": 150, \"water\": 150}\n", x, y)
					break spawn_loop
				}
			}
		}
	} else {
		fmt.Printf("{}\n")
	}
}

func main_ai() {
	fmt.Printf("{}\n")
}

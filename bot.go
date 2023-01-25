package main

import (
	"fmt"
	"golux2/kit"
)

func main() {
	kit.Register(bidder, factory_placer, main_ai)
}

func bidder(gamestate *kit.GameState) {
	fmt.Printf("{\"faction\": \"AlphaStrike\", \"bid\": 0}\n")
}

func factory_placer(gamestate *kit.GameState) {
	fmt.Printf("{}\n")
}

// Note: each time this is called, gamestate is a new object containing new objects, there is no persistence.
func main_ai(gamestate *kit.GameState) {
	fmt.Printf("{}\n")
}

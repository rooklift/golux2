package main

import (
	"fmt"
	"golux2/kit"
)

// First, register your bot with the 3 different functions which will be called at various stages of the game:

func main() {
	kit.Register(bidder, factory_placer, main_ai)
}

// The 3 functions are each called with a Message object which contains almost everything sent by the
// game runner. Note that any objects within are freshly created each time, there is no persistence.

func bidder(msg *kit.Message) {
	fmt.Printf("{\"faction\": \"AlphaStrike\", \"bid\": 0}\n")
}

func factory_placer(msg *kit.Message) {
	fmt.Printf("{}\n")
}

func main_ai(msg *kit.Message) {
	fmt.Printf("{}\n")
}

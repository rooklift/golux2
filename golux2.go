package main

import (
	"golux2/ai"
	"golux2/kit"
)

func main() {
	kit.Run(ai.Bidder, ai.Placer, ai.AI)
}

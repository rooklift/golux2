package kit

import "fmt"

// ------------------------------------------------------------------------------------------------

var bid_string string

func send_bid() {
	fmt.Printf(bid_string)
}

func Bid(faction string, bid int) {
	bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

// ------------------------------------------------------------------------------------------------

var placement_string string

func send_placement() {
	fmt.Printf(placement_string)
}

func PlaceFactory(x int, y int, metal int, water int) {
	placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

// ------------------------------------------------------------------------------------------------

func send_actions() {
	fmt.Printf("{}\n")
}

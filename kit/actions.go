package kit

import "fmt"

var bid_string string
var placement_string string
var factory_actions = make(map[string]int)

func all_action_cleanups() {
	bid_string = "{}\n";
	placement_string = "{}\n";
	for k, _ := range factory_actions {
		delete(factory_actions, k)
	}
}

// ------------------------------------------------------------------------------------------------

func Bid(faction string, bid int) {
	bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func send_bid() {
	fmt.Printf(bid_string)
}

// ------------------------------------------------------------------------------------------------

func PlaceFactory(x int, y int, metal int, water int) {
	placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

func send_placement() {
	fmt.Printf(placement_string)
}

// ------------------------------------------------------------------------------------------------



func FactoryAct(uid string, action int) {
	factory_actions[uid] = action
}

func send_actions() {
	fmt.Printf("{}\n")
}

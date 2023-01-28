package kit

import "fmt"

func (self *Frame) Bid(faction string, bid int) {
	self.bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func (self *Frame) PlaceFactory(x int, y int, metal int, water int) {
	self.placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

func (self *Frame) send_bid() {
	fmt.Printf(self.bid_string)
}

func (self *Frame) send_placement() {
	fmt.Printf(self.placement_string)
}
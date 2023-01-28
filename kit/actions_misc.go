package kit

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (self *Frame) send_actions() {
	var elements []string								// Each element being something like    "factory_0": 1    or    "unit_8": [[0, 1, 0, 0, 0, 1]]
	for key, value := range self.factory_actions {
		elements = append(elements, fmt.Sprintf("\"%s\": %d", key, value))
	}
	for key, value := range self.unit_actions {
		js, err := json.Marshal(value)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		elements = append(elements, fmt.Sprintf("\"%s\": %s", key, js))
	}
	internal := strings.Join(elements, ",")
	fmt.Printf("{" + internal + "}\n")
}

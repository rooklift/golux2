package kit

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (self *Frame) Bid(faction string, bid int) {
	self.bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}", faction, bid)
}

func (self *Frame) PlaceFactory(x int, y int, metal int, water int) {
	self.placement_string = fmt.Sprintf("{\"spawn\": [%d,%d], \"metal\": %d, \"water\": %d}", x, y, metal, water)
}

func (self *Frame) send(s string) {
	if strings.HasSuffix(s, "\n") {									// This would be a programmer error
		panic("send() received already \\n terminated-line")
	}
	fmt.Printf(s)
	fmt.Printf("\n")
	if logging_actions {
		Log("%d > %s", self.RealStep(), s)
	}
}

func (self *Frame) send_bid() {
	self.send(self.bid_string)
}

func (self *Frame) send_placement() {
	self.send(self.placement_string)
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

	self.send("{" + strings.Join(elements, ", ") + "}")

}

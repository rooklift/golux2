package kit

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var frame *Frame
var decoder = json.NewDecoder(os.Stdin)

// Decoders are best for streaming very large lines, I guess. Although the docs claim that a Decoder
// "may read data from r beyond the JSON values requested" it seems that won't happen in practice if
// the thing being read is a whole {}-surrounded object. See https://github.com/golang/go/issues/3942

func Run(bidder func(*Frame), placer func(*Frame), main_ai func(*Frame)) {
	for {
		frame = make_next_frame()
		if frame.Step == 0 {
			bidder(frame)
			frame.send_bid()
		} else if frame.Obs.RealEnvSteps < 0 {
			placer(frame)
			frame.send_placement()
		} else {
			main_ai(frame)
			frame.send_actions()
		}
	}
}

func make_next_frame() *Frame {
	var f *Frame						// Don't try to unmarshal into some already used object since I'm not sure how that works -
	decoder.Decode(&f)					// the rules are complex and in many cases old stuff can persist; see the literature.
	f.clear_actions()
	f.fix_factory_occupancy()
	f.fix_pointers()
	return f
}

func (self *Frame) clear_actions() {
	self.bid_string = "{}\n";
	self.placement_string = "{}\n";
	self.factory_actions = make(map[string]int)
	self.unit_actions = make(map[string][][6]int)
}

func (self *Frame) fix_factory_occupancy() {
	board := self.Obs.Board
	board.FactoryOccupancy = make_2d_int_slice(self.Width(), self.Height(), -1)
	for _, factory := range self.AllFactories() {
		for x := factory.Pos[0] - 1; x <= factory.Pos[0] + 1; x++ {
			for y := factory.Pos[1] - 1; y <= factory.Pos[1] + 1; y++ {
				board.FactoryOccupancy[x][y] = factory.StrainId
			}
		}
	}
}

func (self *Frame) fix_pointers() {
	for _, unit := range self.AllUnits() {
		unit.Frame = self
	}
	for _, factory := range self.AllFactories() {
		factory.Frame = self
	}
}

func (self *Frame) send_bid() {
	fmt.Printf(self.bid_string)
}

func (self *Frame) send_placement() {
	fmt.Printf(self.placement_string)
}

func (self *Frame) send_actions() {
	var elements []string						// Each element being something like    "factory_0": 1    or    "unit_8": [[0, 1, 0, 0, 0, 1]]
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

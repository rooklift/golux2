package kit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------------------------------------------------------------------------------------------

var kframe *Frame						// The kit's frame - not directly accessible by any user calls.
var decoder = json.NewDecoder(os.Stdin)

// Decoders are best for streaming very large lines, I guess. Although the docs claim that a Decoder
// "may read data from r beyond the JSON values requested" it seems that won't happen in practice if
// the thing being read is a whole {}-surrounded object. See https://github.com/golang/go/issues/3942

func Run(bidder func(*Frame), placer func(*Frame), main_ai func(*Frame)) {
	for {
		new_frame := make_next_frame(kframe)
		mangle_frame(kframe)			// Render the old frame unusable so programmers trying to use it notice faster.
		kframe = new_frame
		if kframe.Step == 0 {
			bidder(kframe)
			kframe.send_bid()
		} else if kframe.RealStep() < 0 {
			placer(kframe)
			kframe.send_placement()
		} else {
			main_ai(kframe)
			kframe.send_actions()
		}
	}
}

func make_next_frame(old_frame *Frame) *Frame {

	var f *Frame						// Don't try to unmarshal into some already used object since I'm not sure how that works -
	decoder.Decode(&f)					// the rules are complex and in many cases old stuff can persist; see the literature.

	// Set the action variables to their defaults...

	f.bid_string = "{}";
	f.placement_string = "{}";

	// Create the FactoryOccupancy map...

	f.Obs.Board.FactoryOccupancy = Make2dIntSlice(f.Width(), f.Height(), -1)
	for _, factory := range f.AllFactories() {
		for x := factory.X() - 1; x <= factory.X() + 1; x++ {
			for y := factory.Y() - 1; y <= factory.Y() + 1; y++ {
				f.Obs.Board.FactoryOccupancy[x][y] = factory.StrainId
			}
		}
	}

	// Fix stuff in the units / factories...

	for _, unit := range f.AllUnits() {
		unit.Frame = f
	}
	for _, factory := range f.AllFactories() {
		factory.Frame = f
		factory.Request = -1
	}

	// In the future I might conceivably get main.py to stop sending cfg each turn. We can assume it will not
	// change between turns, so lets just always use the one we got at the start...

	if old_frame != nil {
		f.Info.EnvCfg = old_frame.Info.EnvCfg
	}

	// In the future I might conceivably get main.py to not send valid_spawns_mask once we reach RealStep 0
	// But I guess I'll fill it up with false values just for consistency...

	if len(f.Obs.Board.ValidSpawnsMask) < f.Width() {
		f.Obs.Board.ValidSpawnsMask = Make2dBoolSlice(f.Width(), f.Height(), false)
	}

	// Create cells (technically "nodes") for the A* package...

	return f
}

func mangle_frame(old_frame *Frame) {

	// It's a programmer error to try to reuse the old objects. Let's at least unlink their pointers to the
	// old Frame so it can get garbage collected (also this will cause such bots to error...)

	if old_frame != nil {
		for _, unit := range old_frame.AllUnits() {
			unit.Frame = nil
		}
		for _, factory := range old_frame.AllFactories() {
			factory.Frame = nil
		}
	}
}

// ------------------------------------------------------------------------------------------------

func (self *Frame) send(s string) {
	if strings.HasSuffix(s, "\n") {
		panic("send() received already \\n terminated-line")
	}
	fmt.Printf(s)
	fmt.Printf("\n")
	if logging_actions {
		Log("> %s", s)
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

	for _, factory := range self.MyFactories() {
		if factory.Request >= 0 {
			elements = append(elements, fmt.Sprintf("\"%s\": %d", factory.UnitId, factory.Request))
		}
	}

	for _, unit := range self.MyUnits() {
		if len(unit.Request) == 0 {
			continue
		}
		if len(unit.Request) > 20 {
			Log("%v - attempted to set very long (%d) queue", unit.UnitId, len(unit.Request))
			unit.Request = unit.Request[0:20]
		}
		js, err := json.Marshal(unit.Request)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		if string(js) == "null" {						// Should be impossible as it stands, since we ensure len > 0 above.
			js = []byte("[]")
		}
		elements = append(elements, fmt.Sprintf("\"%s\": %s", unit.UnitId, js))
	}

	self.send("{" + strings.Join(elements, ", ") + "}")

}

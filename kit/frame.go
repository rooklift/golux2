package kit

import (
	"encoding/json"
	"os"
)

var decoder = json.NewDecoder(os.Stdin)

// Decoders are best for streaming very large lines, I guess. Although the docs claim that a Decoder
// "may read data from r beyond the JSON values requested" it seems that won't happen in practice if
// the thing being read is a whole {}-surrounded object. See https://github.com/golang/go/issues/3942

func make_next_frame(old_frame *Frame) *Frame {

	var f *Frame						// Don't try to unmarshal into some already used object since I'm not sure how that works -
	decoder.Decode(&f)					// the rules are complex and in many cases old stuff can persist; see the literature.

	// Set the action variables to their defaults...

	f.bid_string = "{}";
	f.placement_string = "{}";
	f.factory_actions = make(map[string]int)
	f.unit_actions = make(map[string][][6]int)

	// Create the FactoryOccupancy map...

	f.Obs.Board.FactoryOccupancy = make_2d_int_slice(f.Width(), f.Height(), -1)
	for _, factory := range f.AllFactories() {
		for x := factory.Pos[0] - 1; x <= factory.Pos[0] + 1; x++ {
			for y := factory.Pos[1] - 1; y <= factory.Pos[1] + 1; y++ {
				f.Obs.Board.FactoryOccupancy[x][y] = factory.StrainId
			}
		}
	}

	// Fix pointers in the units / factories...

	for _, unit := range f.AllUnits() {
		unit.Frame = f
	}
	for _, factory := range f.AllFactories() {
		factory.Frame = f
	}

	// In the future I might conceivably get main.py to stop sending cfg each turn...

	if old_frame != nil {
		if len(f.Info.EnvCfg) < len(old_frame.Info.EnvCfg) {
			f.Info.EnvCfg = old_frame.Info.EnvCfg
		}
	}

	// In the future I might conceivably get main.py to not send valid_spawns_mask once we reach RealStep 0
	// But I guess I'll fill it up with false values just for consistency...

	if len(f.Obs.Board.ValidSpawnsMask) < f.Width() {
		f.Obs.Board.ValidSpawnsMask = make_2d_bool_slice(f.Width(), f.Height(), false)
	}

	return f
}

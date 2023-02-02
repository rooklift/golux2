package kit

type Frame struct {
	Obs						Obs								`json:"obs"`
	Player					string							`json:"player"`
	RemainingOverageTime	int								`json:"remainingOverageTime"`
	Step					int								`json:"step"`					// Note: different from real_env_steps
	Info					Info							`json:"info"`

	bid_string				string
	placement_string		string
}

// For Info object, see cfg.go

type Obs struct {
	RealEnvSteps			int								`json:"real_env_steps"`
	Board					Board							`json:"board"`			
	Units					map[string]map[string]*Unit		`json:"units"`					// MUST BE POINTER, user sets actions inside the object.
	Factories				map[string]map[string]*Factory	`json:"factories"`				// MUST BE POINTER, user sets actions inside the object.
	Teams					map[string]*Team				`json:"teams"`
}

type Board struct {
	Rubble					[][]int							`json:"rubble"`
	Ice						[][]int							`json:"ice"`
	Ore						[][]int							`json:"ore"`
	Lichen					[][]int							`json:"lichen"`
	LichenStrains			[][]int							`json:"lichen_strains"`
	ValidSpawnsMask			[][]bool						`json:"valid_spawns_mask"`
	FactoryOccupancy		[][]int															// Theoretically part of the Lux API - but generated client side
	FactoriesPerTeam		int								`json:"factories_per_team"`
}

type Unit struct {			// This has a custom unmarshaller, for which see structs_marshal.go
	TeamId					int																// 0 or 1
	UnitId					string															// e.g. "unit_10"
	Power					int
	UnitType				string															// "LIGHT" or "HEAVY"
	Pos
	Cargo
	ActionQueue				[]Action
	
	Frame					*Frame							`json:"-"`
	Request					[]Action
}

type Factory struct {		// This has a custom unmarshaller, for which see structs_marshal.go
	TeamId					int																// 0 or 1
	UnitId					string															// e.g. "factory_4"
	Power					int
	Pos
	Cargo
	StrainId				int																// e.g. 4 - expected to match UnitId
	
	Frame					*Frame							`json:"-"`
	Request					FactoryActionType
}

type Cargo struct {
	Ice						int								`json:"ice"`
	Ore						int								`json:"ore"`
	Water					int								`json:"water"`
	Metal					int								`json:"metal"`
}

type Team struct {
	TeamId					int								`json:"team_id"`				// 0 or 1
	Faction					string							`json:"faction"`
	Water					int								`json:"water"`					// Not updated after factory placement phase!
	Metal					int								`json:"metal"`					// Not updated after factory placement phase!
	FactoriesToPlace		int								`json:"factories_to_place"`
	FactoryStrains			[]int							`json:"factory_strains"`
	PlaceFirst				bool							`json:"place_first"`
	Bid						int								`json:"bid"`
}

type Pos struct {
	X						int
	Y						int
}

type Action struct {
	Type					ActionType
	Direction				Direction
	Resource				Resource
	Amount					int
	Recycle					int
	N						int
}

type ActionType int
type Direction int
type Resource int
type FactoryActionType int

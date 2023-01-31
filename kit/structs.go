package kit

// Note that in various places pointers are used. To my mind, this makes it easier to ensure that
// the user always gets the exact same objects in memory, e.g. when asking for AllUnits(), which
// I think is beneficial...
//
// Also, the way actions work is that, at the end of each turn, the frame looks at its own unit and
// factory objects to see if they are requesting actions, so it MUST be those objects that the user
// applies actions to. 

type Frame struct {
	Obs						*Obs							`json:"obs"`
	Player					string							`json:"player"`
	RemainingOverageTime	int								`json:"remainingOverageTime"`
	Step					int								`json:"step"`					// Note: different from real_env_steps
	Info					*Info							`json:"info"`

	bid_string				string
	placement_string		string
}

type Obs struct {
	RealEnvSteps			int								`json:"real_env_steps"`
	Board					*Board							`json:"board"`			
	Units					map[string]map[string]*Unit		`json:"units"`					// e.g. "player_0" --> "unit_10" --> *Unit
	Factories				map[string]map[string]*Factory	`json:"factories"`				// e.g. "player_0" --> "factory_2" --> *Factory
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

type Unit struct {
	TeamId					int								`json:"team_id"`				// 0 or 1
	UnitId					string							`json:"unit_id"`				// e.g. "unit_10"
	Power					int								`json:"power"`
	UnitType				string							`json:"unit_type"`				// "LIGHT" or "HEAVY"
	Pos						Pos								`json:"pos"`
	Cargo					*Cargo							`json:"cargo"`
	ActionQueue				[]Action						`json:"action_queue"`

	Frame					*Frame							`json:"-"`
	Request					[]Action						`json:"-"`
}

type Factory struct {
	TeamId					int								`json:"team_id"`				// 0 or 1
	UnitId					string							`json:"unit_id"`				// e.g. "factory_4"
	Power					int								`json:"power"`
	Pos						Pos								`json:"pos"`
	Cargo					*Cargo							`json:"cargo"`
	StrainId				int								`json:"strain_id"`				// e.g. 4 - expected to match UnitId

	Frame					*Frame							`json:"-"`
	Request					FactoryActionType				`json:"-"`
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

// ------------------------------------------------------------------------------------------------

type Pos [2]int

type Action struct {
	Type			ActionType
	Direction		Direction
	Resource		Resource
	Amount			int
	Recycle			int
	N				int
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var v [6]int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a.Type = ActionType(v[0])
	a.Direction = Direction(v[1])
	a.Resource = Resource(v[2])
	a.Amount = v[3]
	a.Recycle = v[4]
	a.N = v[5]
	return nil
}

func (a Action) MarshalJSON() ([]byte, error) {
	var v [6]int
	v[0] = int(a.Type)
	v[1] = int(a.Direction)
	v[2] = int(a.Resource)
	v[3] = a.Amount
	v[4] = a.Recycle
	v[5] = a.N
	return json.Marshal(&v)
}

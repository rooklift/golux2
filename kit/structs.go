package kit

import "encoding/json"

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

type Unit struct {
	TeamId					int								// 0 or 1
	UnitId					string							// e.g. "unit_10"
	Power					int
	UnitType				string							// "LIGHT" or "HEAVY"
	Pos
	Cargo
	ActionQueue				[]Action
	
	Frame					*Frame
	Request					[]Action
}

type Factory struct {
	TeamId					int								// 0 or 1
	UnitId					string							// e.g. "factory_4"
	Power					int
	Pos
	Cargo
	StrainId				int								// e.g. 4 - expected to match UnitId
	
	Frame					*Frame
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

// ------------------------------------------------------------------------------------------------
// Thanks to MMJ -- aka themmj on Github -- for help with this...

type Pos struct {
	X						int
	Y						int
}

func (p *Pos) UnmarshalJSON(data []byte) error {
	var v [2]int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	p.X = v[0]
	p.Y = v[1]
	return nil
}

func (p Pos) MarshalJSON() ([]byte, error) {
	var v [2]int
	v[0] = p.X
	v[1] = p.Y
	return json.Marshal(&v)
}

// ------------------------------------------------------------------------------------------------

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

// ------------------------------------------------------------------------------------------------
// Our Factory and Unit types have some embedded structs for user convenience, but this seems to
// mean (as far as I can tell) that we need custom unmarshalling for them... (?)

type unit_tmp struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	UnitType				string							`json:"unit_type"`
	Pos						Pos								`json:"pos"`
	Cargo					Cargo							`json:"cargo"`
	ActionQueue				[]Action						`json:"action_queue"`
}

func (u *Unit) UnmarshalJSON(data []byte) error {

	var v unit_tmp
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	u.TeamId = v.TeamId
	u.UnitId = v.UnitId
	u.Power = v.Power
	u.UnitType = v.UnitType
	u.Pos = v.Pos
	u.Cargo = v.Cargo
	u.ActionQueue = v.ActionQueue

	return nil
}

// ------------------------------------------------------------------------------------------------

type factory_tmp struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	Pos						Pos								`json:"pos"`
	Cargo					Cargo							`json:"cargo"`
	StrainId				int								`json:"strain_id"`
}

func (fc *Factory) UnmarshalJSON(data []byte) error {

	var v factory_tmp
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	fc.TeamId = v.TeamId
	fc.UnitId = v.UnitId
	fc.Power = v.Power
	fc.Pos = v.Pos
	fc.Cargo = v.Cargo
	fc.StrainId = v.StrainId

	return nil
}

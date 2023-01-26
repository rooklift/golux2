package kit

type Message struct {
	Obs						*Obs							`json:"obs"`
	Player					string							`json:"player"`
	RemainingOverageTime	int								`json:"remainingOverageTime"`
	Step					int								`json:"step"`					// Note: different from real_env_steps
	// TODO: info.env_cfg
}

type Obs struct {
	RealEnvSteps			int								`json:"real_env_steps"`
	Board					*Board							`json:"board"`			
	Units					map[string]map[string]*Unit		`json:"units"`
	Factories				map[string]map[string]*Factory	`json:"factories"`
	Teams					map[string]*Team				`json:"teams"`
}

type Board struct {
	Rubble					[][]int							`json:"rubble"`
	Ice						[][]int							`json:"ice"`
	Ore						[][]int							`json:"ore"`
	Lichen					[][]int							`json:"lichen"`
	LichenStrains			[][]int							`json:"lichen_strains"`
	ValidSpawnsMask			[][]bool						`json:"valid_spawns_mask"`
	FactoriesPerTeam		int								`json:"factories_per_team"`
	// TODO: FactoryOccupancy? (which is not given to us)
}

type Unit struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	UnitType				string							`json:"unit_type"`
	Pos						[2]int							`json:"pos"`
	Cargo					*Cargo							`json:"cargo"`
	ActionQueue				[][]int							`json:"action_queue"`
}

type Factory struct {
	TeamId					int								`json:"team_id"`
	UnitId					string							`json:"unit_id"`
	Power					int								`json:"power"`
	Pos						[2]int							`json:"pos"`
	Cargo					*Cargo							`json:"cargo"`
	StrainId				int								`json:"strain_id"`
}

type Cargo struct {
	Ice						int								`json:"ice"`
	Ore						int								`json:"ore"`
	Water					int								`json:"water"`
	Metal					int								`json:"metal"`
}

type Team struct {
	TeamId					int								`json:"team_id"`
	Faction					string							`json:"faction"`
	Water					int								`json:"water"`
	Metal					int								`json:"metal"`
	FactoriesToPlace		int								`json:"factories_to_place"`
	FactoryStrains			[]int							`json:"factory_strains"`
	PlaceFirst				bool							`json:"place_first"`
	Bid						int								`json:"bid"`
}

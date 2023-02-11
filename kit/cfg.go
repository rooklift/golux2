package kit

// Thanks to MMJ -- aka themmj on Github -- for these definitions...

type Info struct {
	EnvCfg							*EnvCfg						`json:"env_cfg"`
}

type EnvCfg struct {
	BiddingSystem					bool						`json:"BIDDING_SYSTEM"`
	CycleLength						int							`json:"CYCLE_LENGTH"`
	DayLength						int							`json:"DAY_LENGTH"`
	FactoryCharge					int							`json:"FACTORY_CHARGE"`
	FactoryProcessingRateMetal		int							`json:"FACTORY_PROCESSING_RATE_METAL"`
	FactoryProcessingRateWater		int							`json:"FACTORY_PROCESSING_RATE_WATER"`
	FactoryRubbleAfterDestruction	int							`json:"FACTORY_RUBBLE_AFTER_DESTRUCTION"`
	FactoryWaterConsumption			int							`json:"FACTORY_WATER_CONSUMPTION"`
	IceWaterRatio					int							`json:"ICE_WATER_RATIO"`
	InitPowerPerFactory				int							`json:"INIT_POWER_PER_FACTORY"`
	InitWaterMetalPerFactory		int							`json:"INIT_WATER_METAL_PER_FACTORY"`
	LichenGainedWithWater			int							`json:"LICHEN_GAINED_WITH_WATER"`
	LichenLostWithoutWater			int							`json:"LICHEN_LOST_WITHOUT_WATER"`
	LichenWateringCostFactor		float64						`json:"LICHEN_WATERING_COST_FACTOR"`
	MaxLichenPerTile				int							`json:"MAX_LICHEN_PER_TILE"`
	MaxFactories					int							`json:"MAX_FACTORIES"`
	MaxRubble						int							`json:"MAX_RUBBLE"`
	MinFactories					int							`json:"MIN_FACTORIES"`
	MinLichenToSpread				int							`json:"MIN_LICHEN_TO_SPREAD"`
	OreMetalRatio					int							`json:"ORE_METAL_RATIO"`
	PowerLossFactor					float64						`json:"POWER_LOSS_FACTOR"`
	PowerPerConnectedLichenTile		int							`json:"POWER_PER_CONNECTED_LICHEN_TILE"`
	Robots							map[string]*UnitConfig		`json:"ROBOTS"`
	UnitActionQueueSize				int							`json:"UNIT_ACTION_QUEUE_SIZE"`
	MapSize							int							`json:"map_size"`
	MaxEpisodeLength				int							`json:"max_episode_length"`
	MaxTransferAmount				int							`json:"max_transfer_amount"`
	ValidateActionSpace				bool						`json:"validate_action_space"`
	Verbose							int							`json:"verbose"`
}

type UnitConfig struct {
	ActionQueuePowerCost			int							`json:"ACTION_QUEUE_POWER_COST"`
	BatteryCapacity					int							`json:"BATTERY_CAPACITY"`
	CargoSpace						int							`json:"CARGO_SPACE"`
	Charge							int							`json:"CHARGE"`
	DigCost							int							`json:"DIG_COST"`
	DigLichenRemoved				int							`json:"DIG_LICHEN_REMOVED"`
	DigResourceGain					int							`json:"DIG_RESOURCE_GAIN"`
	DigRubbleRemoved				int							`json:"DIG_RUBBLE_REMOVED"`
	InitPower						int							`json:"INIT_POWER"`
	MetalCost						int							`json:"METAL_COST"`
	MoveCost						float64						`json:"MOVE_COST"`
	PowerCost						int							`json:"POWER_COST"`
	RubbleAfterDestruction			int							`json:"RUBBLE_AFTER_DESTRUCTION"`
	RubbleMovementCost				float64						`json:"RUBBLE_MOVEMENT_COST"`
	SelfDestructCost				int							`json:"SELF_DESTRUCT_COST"`
}

func (self *Frame) GetCfg() *EnvCfg {
	return self.Info.EnvCfg
}

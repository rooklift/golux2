package kit

const (

	MOVE			ActionType = 0
	TRANSFER		ActionType = 1
	PICKUP			ActionType = 2
	DIG				ActionType = 3
	SELFDESTRUCT	ActionType = 4
	RECHARGE		ActionType = 5

	CENTER			Direction = 0
	CENTRE			Direction = 0		// The true spelling
	UP				Direction = 1
	NORTH			Direction = 1
	RIGHT			Direction = 2
	EAST			Direction = 2
	DOWN			Direction = 3
	SOUTH			Direction = 3
	LEFT			Direction = 4
	WEST			Direction = 4

	ICE				Resource = 0
	ORE				Resource = 1
	WATER			Resource = 2
	METAL			Resource = 3
	POWER			Resource = 4

	BUILD_LIGHT		FactoryActionType = 0
	BUILD_HEAVY		FactoryActionType = 1
	WATER_LICHEN	FactoryActionType = 2

)

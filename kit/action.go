package kit

// Thanks to MMJ -- aka themmj on Github

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

// [0]     [1]        [2]       [3]     [4]      [5]
// ACTION  DIRECTION  RESOURCE  AMOUNT  RECYCLE  N

func MoveAction(direction Direction, recycle int, n int) Action {
	return Action{MOVE, direction, 0, 0, recycle, n}
}

func TransferAction(direction Direction, resource Resource, amount int, recycle int, n int) Action {
	return Action{TRANSFER, direction, resource, amount, recycle, n}
}

func PickupAction(resource Resource, amount int, recycle int, n int) Action {
	return Action{PICKUP, 0, resource, amount, recycle, n}
}

func DigAction(recycle int, n int) Action {
	return Action{DIG, 0, 0, 0, recycle, n}
}

func SelfDestructAction() Action {
	return Action{SELFDESTRUCT, 0, 0, 0, 0, 1}
}

func RechargeAction(amount int, recycle int, n int) Action {
	return Action{RECHARGE, 0, 0, amount, recycle, n}
}

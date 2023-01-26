package kit

import (
	"encoding/json"
	"fmt"
	"strings"
)

var bid_string string
var placement_string string
var factory_actions = make(map[string]int)
var robot_actions = make(map[string][][]int)

func all_action_cleanups() {
	bid_string = "{}\n";
	placement_string = "{}\n";
	for k, _ := range factory_actions {
		delete(factory_actions, k)
	}
	for k, _ := range robot_actions {
		delete(robot_actions, k)
	}
}

// ------------------------------------------------------------------------------------------------

func Bid(faction string, bid int) {
	bid_string = fmt.Sprintf("{\"faction\": \"%s\", \"bid\": %d}\n", faction, bid)
}

func send_bid() {
	fmt.Printf(bid_string)
}

// ------------------------------------------------------------------------------------------------

func PlaceFactory(x int, y int, metal int, water int) {
	placement_string = fmt.Sprintf("{\"spawn\": [%d, %d], \"metal\": %d, \"water\": %d}\n", x, y, metal, water)
}

func send_placement() {
	fmt.Printf(placement_string)
}

// ------------------------------------------------------------------------------------------------

func FactoryAct(uid string, action int) {
	factory_actions[uid] = action
}

func RobotAct(uid string, action [][]int) {
	robot_actions[uid] = action
}

func send_actions() {

	var elements []string				// Each element being something like    "factory_0": 1    or    "unit_8": [[0, 1, 0, 0, 0, 1]]

	for key, value := range factory_actions {
		elements = append(elements, fmt.Sprintf("\"%s\": %d", key, value))
	}

	for key, value := range robot_actions {
		js, err := json.Marshal(value)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		elements = append(elements, fmt.Sprintf("\"%s\": %s", key, js))
	}

	internal := strings.Join(elements, ",")
	fmt.Printf("{" + internal + "}\n")
}

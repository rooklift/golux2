package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var msg *Message

var scanner = bufio.NewScanner(os.Stdin)

var bid_string string
var placement_string string
var factory_actions = make(map[string]int)
var robot_actions = make(map[string][][6]int)

func Run(bidder func(), placer func(), main_ai func()) {
	for {
		update()
		if msg.Step == 0 {
			bidder()
			send_bid()
		} else if msg.Obs.RealEnvSteps < 0 {
			placer()
			send_placement()
		} else {
			main_ai()
			send_actions()
		}
	}
}

func update() {

	all_action_cleanups()

	var new_msg *Message						// Don't try to unmarshal into the already extant message since I'm not sure how that works -
												// the rules are complex and in many cases old objects can persist; see the literature.
	scanner.Scan()
	err := json.Unmarshal(scanner.Bytes(), &new_msg)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	msg = new_msg
}

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

func send_bid() {
	fmt.Printf(bid_string)
}

func send_placement() {
	fmt.Printf(placement_string)
}

func send_actions() {
	var elements []string						// Each element being something like    "factory_0": 1    or    "unit_8": [[0, 1, 0, 0, 0, 1]]
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

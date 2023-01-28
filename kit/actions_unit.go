package kit

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (self *Unit) SetQueue(args ...[6]int) {
	var queue [][6]int
	for _, item := range args {
		queue = append(queue, item)
	}
	self.Frame.unit_actions[self.UnitId] = queue
}

func (self *Unit) Cancel() {
	delete(self.Frame.unit_actions, self.UnitId)
}

func Action(atype int, direction int, resource int, amount int, send_to_back bool, iterations int) [6]int {
	return [6]int{atype, direction, resource, amount, bool_to_int(send_to_back), iterations}
}

func (self *Frame) send_actions() {
	var elements []string								// Each element being something like    "factory_0": 1    or    "unit_8": [[0, 1, 0, 0, 0, 1]]
	for key, value := range self.factory_actions {
		elements = append(elements, fmt.Sprintf("\"%s\": %d", key, value))
	}
	for key, value := range self.unit_actions {
		js, err := json.Marshal(value)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		elements = append(elements, fmt.Sprintf("\"%s\": %s", key, js))
	}
	internal := strings.Join(elements, ",")
	fmt.Printf("{" + internal + "}\n")
}

package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var logfile, _ = os.Create("log.txt")

var msg *Message

func Run(bidder func(), placer func(), mainai func()) {
	for {
		update()
		if msg.Step == 0 {
			bidder()
			send_bid()
		} else if msg.Obs.RealEnvSteps < 0 {
			placer()
			send_placement()
		} else {
			mainai()
			send_actions()
		}
	}
}

func update() {

	all_action_cleanups()

	var new_msg *Message						// Don't try to unmarshal into the already extant message since I'm not sure how that works -
												// the rules are complex and in many cases old objects can persist; see the literature.
	scanner.Scan()

	logfile.Write(scanner.Bytes())
	logfile.Write([]byte("\n"))

	err := json.Unmarshal(scanner.Bytes(), &new_msg)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	msg = new_msg
}

package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var msg *Message

var scanner = bufio.NewScanner(os.Stdin)
var logfile, _ = os.Create("log.txt")

func Register(bidder func(), placer func(), mainai func()) {

	update()
	bidder()
	fmt.Printf(bid_string)

	for {
		update()
		if msg.Obs.RealEnvSteps < 0 {
			placer()
			fmt.Printf(placement_string)
		} else {
			mainai()							// Call main AI
		}
	}
}

func update() {

	bid_string = ""								// Clear any user-created messages from earlier...
	placement_string = ""

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

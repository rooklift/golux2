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

func Register(fn1 func(), fn2 func(), fn3 func()) {
	read_update()
	fn1()										// Call bidder
	for {
		read_update()
		if msg.Obs.RealEnvSteps < 0 {
			fn2()								// Call factory placer
		} else {
			fn3()								// Call main AI
		}
	}
}

func read_update() {					
	var new_msg *Message						// Don't try to unmarshal into the already extant message since I'm not sure how that works -
	scanner.Scan()								// the rules are complex and in many cases old objects can persist; see the literature.
	logfile.Write(scanner.Bytes())
	logfile.Write([]byte("\n"))
	err := json.Unmarshal(scanner.Bytes(), &new_msg)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	msg = new_msg
}

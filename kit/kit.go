package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var last_message *Message

var scanner = bufio.NewScanner(os.Stdin)
var logfile, _ = os.Create("log.txt")

func Register(fn1 func(*Message), fn2 func(*Message), fn3 func(*Message)) {
	read_update()
	fn1(last_message)								// Call bidder
	for {
		read_update()
		if last_message.Obs.RealEnvSteps < 0 {
			fn2(last_message)						// Call factory placer
		} else {
			fn3(last_message)						// Call main AI
		}
	}
}

func read_update() {					
	var message Message								// Don't try to unmarshal into the already extant last_message since I'm not sure how that works -
	scanner.Scan()									// in certain situations (maybe not this one) such an act can allow old objects to persist, don't want to worry.
	logfile.Write(scanner.Bytes())
	logfile.Write([]byte("\n"))
	err := json.Unmarshal(scanner.Bytes(), &message)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	last_message = &message
}

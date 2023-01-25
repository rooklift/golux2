package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

var gamestate GameState

func Register() {
	read_update()
	fmt.Printf("{\"faction\": \"AlphaStrike\", \"bid\": 0}\n")
	for {
		read_update()
		fmt.Printf("{}\n")
	}
}

func read_update() {

	var message Message1

	scanner.Scan()

	err := json.Unmarshal(scanner.Bytes(), &message)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	gamestate = message.GameState

}

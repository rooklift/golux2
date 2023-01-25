package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

var gamestate GameState

var logfile, _ = os.Create("log.txt")

func Register() {
	read_line_one()
	fmt.Printf("{\"faction\": \"Golang\", \"bid\": 0}\n")
	for {
		read_update()
		fmt.Printf("{}\n")
	}
}

func read_line_one() {

	var message Message1

	scanner.Scan()

	logfile.Write(scanner.Bytes())

	err := json.Unmarshal(scanner.Bytes(), &message)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	gamestate = message.GameState

}

func read_update() {

	var message Message2			// Allegedly, the agent might actually receive the full gamestate, so this may be wrong.

	scanner.Scan()

	err := json.Unmarshal(scanner.Bytes(), &message)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	update := message.GameStateUpdate

	gamestate.EnvSteps = update.EnvSteps
	gamestate.Units = update.Units
	gamestate.Factories = update.Factories
	gamestate.Teams = update.Teams

	// TODO - parse update.Board

}

package kit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

var gamestate *GameState

func Register(fn1 func(*GameState), fn2 func(*GameState), fn3 func(*GameState)) {
	read_update()
	fn1(gamestate)
	for {
		read_update()
		if gamestate.EnvSteps < 0 {
			fn2(gamestate)
		} else {
			fn3(gamestate)
		}
	}
}

func read_update() {
	var message Message1
	scanner.Scan()
	err := json.Unmarshal(scanner.Bytes(), &message)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	gamestate = &message.GameState
}

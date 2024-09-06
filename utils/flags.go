package utils

import (
	"errors"
	"flag"
	"fmt"

	c "github.com/kelbwah/go-chess/constants"
)

func ParseFlag() (c.Runtime, error) {
	serverFlag := flag.Bool("server", false, "Start in server mode")
	gameFlag := flag.Bool("game", false, "Start in game mode")

	flag.Parse()

	if *serverFlag {
		fmt.Println("Initiating server...")
		return c.SERVER, nil
	} else if *gameFlag {
		fmt.Println("Loading game...")
		return c.GAME, nil
	} else {
		return "", errors.New("Please specify if you want to run a '--server' or a '--game'.")
	}
}

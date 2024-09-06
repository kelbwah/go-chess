package main

import (
	"fmt"
	"os"

	c "github.com/kelbwah/go-chess/constants"
	"github.com/kelbwah/go-chess/utils"
)

func main() {
	flag, err := utils.ParseFlag()
	if err != nil {
		fmt.Printf("Error while parsing flags: %s.", err.Error())
		os.Exit(1)
	}

	if flag == c.GAME {
		utils.InitDrawLoop()
	} else {
		fmt.Println("server!")
	}
}

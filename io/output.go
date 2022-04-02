package io

import (
	"log"
	"os"

	"github.com/owsky/game-libraries-crosschecker/steamapi"
)

func PrintOutput(matchedGames []steamapi.Game) {
	if file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777); err == nil {
		defer file.Close()
		// writes to the file each entry of the matched games
		for _, game := range matchedGames {
			file.WriteString(game.Name + "\n")
		}
		log.Println("Written output to output.txt")
	} else {
		log.Fatalln(err)
	}
}

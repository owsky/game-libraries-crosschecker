package io

import (
	"log"
	"os"

	"github.com/owsky/game-libraries-crosschecker/steamapi"
)

func PrintOutput(matchedGames []steamapi.Game) {
	// os.Remove("output.txt")
	// file, err := os.Create("output.txt")
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	// writes to the file each entry of the matched games
	for _, game := range matchedGames {
		file.WriteString(game.Name + "\n")
	}
	defer file.Close()
}

package main

import (
	"log"
	"os"

	"github.com/owsky/game-libraries-crosschecker/checker"
	"github.com/owsky/game-libraries-crosschecker/io"
	repo "github.com/owsky/game-libraries-crosschecker/repository"
)

func main() {
	// retrieves the whole list of games available on Steam
	allSteamGames := repo.RequestAllSteamGames()
	// retrieves the list of games owned by the user
	ownedSteamGames := repo.RequestOwnedSteamGames()
	// consumes the CSV file exported by PlayNite
	thirdPartyLibraries := io.Ingester()
	// returns the previous list sans the games not available on Steam and the ones already owned on Steam
	matchedGames := checker.Crosscheck(thirdPartyLibraries, allSteamGames, ownedSteamGames)

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

package main

import (
	"log"

	"github.com/owsky/game-libraries-crosschecker/checker"
	"github.com/owsky/game-libraries-crosschecker/io"
	"github.com/owsky/game-libraries-crosschecker/steamapi"
)

func main() {
	// retrieves the whole list of games available on Steam
	allSteamGames := steamapi.GetWholeSteamLibrary()
	// retrieves the list of games owned by the user
	ownedSteamGames := steamapi.GetOwnSteamLibrary()
	// consumes the CSV file exported by PlayNite
	thirdPartyLibraries := io.Ingester()
	// returns the previous list sans the games not available on Steam and the ones already owned on Steam
	matchedGames := checker.Crosscheck(thirdPartyLibraries, allSteamGames, ownedSteamGames)

	log.Println(thirdPartyLibraries[:10])

	io.PrintOutput(matchedGames)

	io.Ingester()
}

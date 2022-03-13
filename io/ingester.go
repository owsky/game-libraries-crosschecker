package io

import (
	"log"
	"os"
	"sort"

	"github.com/gocarina/gocsv"
)

// removes all Steam games from the list provided by PlayNite
func purgeSteam(games []*GameIngester) []string {
	var gameNames []string
	for _, game := range games {
		if game.Source != "Steam" {
			gameNames = append(gameNames, game.Name)
		}
	}
	return gameNames
}

// this library only supports csv files formatted as standard UTF-8, files formatted in UTF-8-BOM will silently fail
// playnite also adds quotes to the column headers which is out of spec
// I'm planning to address this issue soon
func Ingester() []string {
	gamesFile, err := os.Open("games.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer gamesFile.Close()

	games := []*GameIngester{}

	if err := gocsv.UnmarshalFile(gamesFile, &games); err != nil {
		log.Fatalln(err)
	}

	sort.Slice(games, func(i, j int) bool {
		return games[i].Name < games[j].Name
	})

	return purgeSteam(games)
}

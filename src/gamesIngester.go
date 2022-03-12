package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Game struct {
	Name        string `csv:"Name"`
	Source      string `csv:"Source"`
	ReleaseDate string `csv:"ReleaseDate"`
	PlayTime    string `csv:"Playtime"`
	IsInstalled string `csv:"IsInstalled"`
}

func purgeSteam(games []*Game) []string {
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
func Ingester() []string {
	// gamesFile, err := os.OpenFile("games.csv", os.O_RDONLY, os.ModePerm)
	gamesFile, err := os.Open("games.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer gamesFile.Close()

	games := []*Game{}

	if err := gocsv.UnmarshalFile(gamesFile, &games); err != nil {
		log.Fatalln(err)
	}

	return purgeSteam(games)
}

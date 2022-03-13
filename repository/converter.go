package repository

import (
	"github.com/owsky/game-libraries-crosschecker/repository/steamapi"
)

// convert from the deserialized structs to the generic game struct

func appToGame(games []steamapi.App) []Game {
	var result []Game
	for _, g := range games {
		result = append(result, Game{Name: g.Name, AppID: int(g.Appid)})
	}
	return result
}

func ownedToGame(games []steamapi.OwnedGame) []Game {
	var result []Game
	for _, g := range games {
		result = append(result, Game{Name: g.Name, AppID: int(g.AppID)})
	}
	return result
}

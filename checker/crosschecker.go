package checker

import (
	"errors"

	"github.com/owsky/game-libraries-crosschecker/steamapi"
)

// binary search of the provided game's name through a list of games
func contains(apps []steamapi.Game, name string) (steamapi.Game, error) {
	left := 0
	right := len(apps) - 1

	for left <= right {
		mid := (left + right) / 2

		if apps[mid].Name == name {
			return apps[mid], nil
		} else if apps[mid].Name < name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	err := errors.New("app not found")
	return steamapi.Game{}, err
}

func Crosscheck(games []string, allSteamGames []steamapi.Game, ownedSteamGames []steamapi.Game) []steamapi.Game {
	var result []steamapi.Game
	for _, game := range games {
		// checks whether the provided game is available on Steam
		app, err := contains(allSteamGames, game)
		if err == nil {
			// checks whether the user doesn't already own the game
			_, err := contains(ownedSteamGames, game)
			if err != nil {
				result = append(result, app)
			}
		}
	}
	return result
}

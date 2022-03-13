package repository

import (
	"github.com/owsky/game-libraries-crosschecker/repository/steamapi"
)

// wrapper methods through the single source of truth

func RequestAllSteamGames() []Game {
	res := steamapi.GetAppList()
	return appToGame(res)
}

func RequestOwnedSteamGames() []Game {
	res := steamapi.GetOwnedGames()
	return ownedToGame(res)
}

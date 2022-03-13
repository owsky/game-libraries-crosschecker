package steamapi

type OwnedGame struct {
	Name  string `json:"name"`
	AppID int    `json:"appid"`
}

type OwnedGames struct {
	GameCount int         `json:"game_count"`
	Games     []OwnedGame `json:"games"`
}

type OwnedGamesResponse struct {
	Response OwnedGames `json:"response"`
}

package steamapi

type OwnedGames struct {
	GameCount int    `json:"game_count"`
	Games     []Game `json:"games"`
}

type OwnedGamesResponse struct {
	Response OwnedGames `json:"response"`
}

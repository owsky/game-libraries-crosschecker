package config

type Config struct {
	CsvFilePath string `json:"csvFilePath"`
	SteamUID    string `json:"steamUID"`
	SteamAPIKey string `json:"steamAPIKey"`
}

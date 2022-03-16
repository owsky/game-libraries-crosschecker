package steamapi

type AppList struct {
	Apps []Game `json:"apps"`
}

type SteamResult struct {
	SteamResult AppList `json:"applist"`
}

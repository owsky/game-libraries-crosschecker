package main

type App struct {
	Appid uint32 `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

type SteamResult struct {
	SteamResult AppList `json:"applist"`
}

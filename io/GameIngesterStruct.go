package io

type GameIngester struct {
	Name        string `csv:"Name"`
	Source      string `csv:"Source"`
	ReleaseDate string `csv:"ReleaseDate"`
	PlayTime    string `csv:"Playtime"`
	IsInstalled string `csv:"IsInstalled"`
}

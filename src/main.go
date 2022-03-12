package main

import (
	"log"
	"os"
)

func main() {
	gamesList := GetAppList()
	games := Ingester()

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for _, game := range games {
		g, err := Contains(gamesList, game)
		if err == nil {
			file.WriteString(g.Name + "\n")
		}
	}
}

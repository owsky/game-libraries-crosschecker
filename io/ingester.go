package io

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func Ingester() []GameIngester {
	file, err := os.Open("games.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var res []GameIngester
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawLine := scanner.Text()

		if !strings.Contains(rawLine, "#TYPE") && !strings.Contains(rawLine, "\"Name\",\"Source\",\"ReleaseDate\",\"Playtime\",\"IsInstalled\"") {
			line := strings.Split(rawLine, ",")
			if len(line) == 5 && !strings.Contains(line[1], "Steam") {
				res = append(res, GameIngester{
					Name:        line[0],
					Source:      line[1],
					ReleaseDate: line[2],
					PlayTime:    line[3],
					IsInstalled: line[4],
				})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})

	return res
}

package steamapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

// sends a get request to the steam api to retrieve the list of games owned by the user
func GetOwnSteamLibrary() []Game {
	steamUID := os.Getenv("uid")
	steamAPIKey := os.Getenv("api")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/", nil)
	if err != nil {
		log.Fatalln(err)
	}
	q := request.URL.Query()
	q.Add("key", steamAPIKey)
	q.Add("steamid", steamUID)
	q.Add("include_appinfo", "true")
	q.Add("include_played_free_games", "true")
	request.URL.RawQuery = q.Encode()

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data OwnedGamesResponse
	unmarshallErr := json.Unmarshal(body, &data)
	if unmarshallErr != nil {
		log.Fatalln("There was an error retrieving your Steam library. Check input data")
	}

	sort.Slice(data.Response.Games, func(i, j int) bool {
		return data.Response.Games[i].Name < data.Response.Games[j].Name
	})

	return data.Response.Games
}

package steamapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/owsky/game-libraries-crosschecker/config"
)

// sends a get request to the steam api to retrieve the list of games owned by the user
func GetOwnedGames() []OwnedGame {
	config := config.GetConfigVariables()

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/", nil)
	q := request.URL.Query()
	q.Add("key", config.ApiKey)
	q.Add("steamid", config.SteamID)
	q.Add("include_appinfo", "true")
	q.Add("include_played_free_games", "true")
	request.URL.RawQuery = q.Encode()

	if err != nil {
		log.Fatalln(err)
	}

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
		log.Fatalln(err)
	}

	sort.Slice(data.Response.Games, func(i, j int) bool {
		return data.Response.Games[i].Name < data.Response.Games[j].Name
	})

	return data.Response.Games
}

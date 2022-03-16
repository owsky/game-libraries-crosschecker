package steamapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// sends a get request to the steam api to retrieve the whole list of games available
func GetWholeSteamLibrary() []Game {
	reqBody, reqErr := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	jsonBody, jsonErr := ioutil.ReadAll(reqBody.Body)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	var res SteamResult
	unmarshallErr := json.Unmarshal(jsonBody, &res)
	if unmarshallErr != nil {
		log.Fatalln(unmarshallErr)
	}
	if len(res.SteamResult.Apps) == 0 {
		err := errors.New("something went wrong while retrieving the games from the steam api")
		log.Fatalln(err)
	}

	sort.Slice(res.SteamResult.Apps, func(i, j int) bool {
		return res.SteamResult.Apps[i].Name < res.SteamResult.Apps[j].Name
	})

	return res.SteamResult.Apps
}

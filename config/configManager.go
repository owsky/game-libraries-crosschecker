package config

import (
	"encoding/json"
	"log"
	"os"
)

// retrieves the configuration variables from config.json file
// located in the same directory of the executable
// the expected variables are:
// steamID: containing the ID of the user in uint64 format
// apiKey: the API key of the user querying the steam API
func GetConfigVariables() Config {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	var jsonBody Config
	jsonErr := json.Unmarshal(file, &jsonBody)
	if jsonErr != nil {
		log.Fatalln(err)
	}

	return jsonBody
}

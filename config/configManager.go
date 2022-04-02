package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func createJsonFile() {
	config := Config{
		CsvFilePath: "",
		SteamUID:    "",
		SteamAPIKey: "",
	}

	file, _ := json.MarshalIndent(config, "", "")
	ioutil.WriteFile("config.json", file, 0777)
	log.Fatalln("Config file has been created. Insert values and execute again")
}

func SetEnvironmentVariables(args []string) {
	var csvFilePath string
	var steamUID string
	var steamAPIKey string

	if len(args) > 0 {
		if len(args) < 3 {
			log.Println("The number of provided arguments was incorrect")
			log.Println("The correct parameters are: <csv-file-path> <steam-uid> <steam-api-key>")
			log.Println("You can also provide a config.json file, a preformatted file will be created if no proper config is found when run without cli arguments")
			os.Exit(1)
		}

		csvFilePath = args[0]
		steamUID = args[1]
		steamAPIKey = args[2]
	} else {
		if f, err := os.OpenFile("config.json", os.O_RDONLY, 0777); err == nil {
			defer f.Close()

			if byteValue, err := ioutil.ReadAll(f); err == nil {
				var config Config
				if unmarshallErr := json.Unmarshal(byteValue, &config); unmarshallErr == nil {
					csvFilePath = config.CsvFilePath
					steamUID = config.SteamUID
					steamAPIKey = config.SteamAPIKey
				}
			}
		} else {
			createJsonFile()
		}
	}

	if len(csvFilePath) > 0 && len(steamUID) > 0 && len(steamAPIKey) > 0 {
		os.Setenv("csv", csvFilePath)
		os.Setenv("uid", steamUID)
		os.Setenv("api", steamAPIKey)
	} else {
		log.Fatalln("One or more fields are empty in config file")
	}
}

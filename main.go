package main

import (
	"fmt"
	"log"

	"github.com/Conor-Flemign/CraigsFinder/craigslist"
)

func main() {
	configFile := "config.yaml"

	//load our config file and parse the values to our structs
	cfg, err := loadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return

	}

	//populate map with valid search areas
	areasMap := make(map[string]craigslist.Area)
	areas, err := craigslist.FetchAreas(cfg.APIURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, a := range areas {
		areasMap[a.Hostname] = a
	}

	for _, search := range cfg.Searches {
		if _, ok := areasMap[search.Area]; !ok {
			log.Printf("Area '%s' is not available for search. Skipping...", search.Area)
			continue
		}
	}
}

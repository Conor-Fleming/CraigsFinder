package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Conor-Fleming/CraigsFinder/craigslist"
)

const configFile string = "config.yaml"

func main() {
	//Log setup
	f, err := os.OpenFile("ErrorLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Print("This log will contain any error messages encountered while running\n\n")

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

	//verifying the configured search locations are valid craigslist locations
	searchesToRun := make([]craigslist.Search, 0)
	for _, search := range cfg.Searches {
		if _, ok := areasMap[search.Location]; !ok {
			log.Printf("Area '%s' is not available for search. Skipping...", search.Location)
			continue
		}

		//populate area field of search object and add to slice
		search.Area = areasMap[search.Location]
		searchesToRun = append(searchesToRun, search)
	}

	//spin off go routines to run each search
	for _, search := range searchesToRun {
		go craigslist.RunSearch(search)
	}
}

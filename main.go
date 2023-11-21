package main

import (
	"fmt"
	"log"
	"os"

	cl "github.com/Conor-Flemign/CraigsFinder/craigslist/"
)

const configFile string = "config.yaml"

func main() {
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//load our config file and parse the values to our structs
	cfg, err := loadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return

	}

	//populate map with valid search areas
	areasMap := make(map[string]cl.Area)
	areas, err := cl.FetchAreas(cfg.APIURL)
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

		area := areasMap[search.Area]

		fmt.Println(area, search)
		//format data
		//run search
		//
	}
}

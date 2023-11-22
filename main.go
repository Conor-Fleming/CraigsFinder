package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Conor-Fleming/CraigsFinder/config"
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
	cfg, err := config.LoadConfig(configFile)
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
	searchesToRun := verifySearchParams(cfg, areasMap)

	//setup search logs and defer closure
	for _, v := range searchesToRun {
		initLogger(v.Term)
	}
	defer closeLoggers()

	//setup logs & spin off go routines to run each search
	for _, search := range searchesToRun {
		go craigslist.RunSearch(search)
	}
}

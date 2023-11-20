package main

import (
	"flag"
	"fmt"

	"github.com/Conor-Flemign/CraigsFinder/craigslist"
)

func main() {
	configFile := "config.yml"

	cfg, err := loadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Define command line flags
	hostnamePtr := flag.String("hostname", cfg.DefaultHostname, "Hostname for craigslist area details")
	flag.Parse()

	areas, err := craigslist.FetchAreas(cfg.APIURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	areasMap := make(map[string]craigslist.Area)
	for _, a := range areas {
		areasMap[a.Hostname] = a
	}

	// Accessing values in the map by provided or default hostname
	if area, ok := areasMap[*hostnamePtr]; ok {
		fmt.Printf("Area: %s, ID: %d, Latitude: %f, Longitude: %f\n", area.Description, area.AreaID, area.Latitude, area.Longitude)
	} else {
		fmt.Println("Area not found for provided hostname. Using default values.")
		defaultArea := areasMap[cfg.DefaultHostname]
		fmt.Printf("Default Area: %s, ID: %d, Latitude: %f, Longitude: %f\n", defaultArea.Description, defaultArea.AreaID, defaultArea.Latitude, defaultArea.Longitude)
	}
}

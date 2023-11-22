package main

import (
	"fmt"
	"log"

	"github.com/Conor-Fleming/CraigsFinder/config"
	"github.com/Conor-Fleming/CraigsFinder/craigslist"
)

func verifySearchParams(cfg config.Config, areasMap map[string]craigslist.Area) []craigslist.Search {
	searchesToRun := make([]craigslist.Search, 0)
	for _, search := range cfg.Searches {
		if _, ok := areasMap[search.Location]; !ok {
			log.Printf("Area '%s' is not available for search. Skipping...", search.Location)
			continue
		}
		fmt.Println(search)
		//populate area field of search object and add to slice
		search.Region = areasMap[search.Location]
		searchesToRun = append(searchesToRun, search)
	}

	fmt.Println("Here")
	return searchesToRun
}

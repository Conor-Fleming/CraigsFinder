package main

import (
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

		//populate area field of search object and add to slice
		search.Area = areasMap[search.Location]
		searchesToRun = append(searchesToRun, search)
	}

	return searchesToRun
}

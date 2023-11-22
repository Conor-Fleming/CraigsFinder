package craigslist

import (
	"fmt"
	"log"
	"time"

	"github.com/Conor-Fleming/CraigsFinder/logs"
	"github.com/ecnepsnai/craigslist"
)

type Search struct {
	Location string `yaml:"area"`
	Category string `yaml:"category"`
	Term     string `yaml:"term"`
	Region   Area
}

func RunSearch(search Search) {
	logs.LogTo(search.Term, fmt.Sprintf("Starting search for %s", search.Term))
	results, err := craigslist.Search(search.Category, search.Term, craigslist.LocationParams{
		AreaID:         search.Region.AreaID,
		Latitude:       search.Region.Latitude,
		Longitude:      search.Region.Longitude,
		SearchDistance: 30,
	})

	if err != nil {
		log.Printf("there was an error running the search: %v", err)
	}

	if len(results) == 0 {
		logs.LogTo(search.Term, fmt.Sprintf("Running Search %s --- No results!", time.Now()))
	}

	for _, v := range results {
		posting, err := v.Posting()
		if err != nil {
			log.Printf("there was an error getting the posting details: %v", err)
			continue
		}

		fmt.Printf("%s - $%v - %v\n", posting.Title, posting.Price, posting.URL)
		fmt.Printf("--------\n")
	}
}

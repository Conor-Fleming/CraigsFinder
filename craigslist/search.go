package craigslist

import (
	"fmt"
	"log"

	"github.com/ecnepsnai/craigslist"
)

type Search struct {
	Location string `yaml:"area"`
	Category string `yaml:"category"`
	Term     string `yaml:"term"`
	Area     Area
}

func RunSearch(search Search) {
	log.Printf("Running search on:%v, %v, %v", search.Area.AreaID, search.Category, search.Term)

	results, err := craigslist.Search(search.Category, search.Term, craigslist.LocationParams{
		AreaID:         search.Area.AreaID,
		Latitude:       search.Area.Latitude,
		Longitude:      search.Area.Longitude,
		SearchDistance: 30,
	})

	if err != nil {
		log.Printf("there was an error running the search: %v", err)
	}

	if len(results) == 0 {
		log.Printf("Running Search %s --- No results!", search.Term)
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

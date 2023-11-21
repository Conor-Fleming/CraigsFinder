package craigslist

import (
	"fmt"
	"log"

	"github.com/ecnepsnai/craigslist"
)

type Search struct {
	Area     string `yaml:"area"`
	Category string `yaml:"category"`
	Term     string `yaml:"term"`
}

type Config struct {
	Searches []Search `yaml:"searches"`
	APIURL   string   `yaml:"api_url"`
}

func RunSearch(area Area, search Search) {
	results, err := craigslist.Search(search.Category, search.Term, craigslist.LocationParams{
		AreaID:         area.AreaID,
		Latitude:       area.Latitude,
		Longitude:      area.Longitude,
		SearchDistance: 30,
	})

	if err != nil {
		log.Printf("there was an error running the search: %v", err)
	}

	if len(results) == 0 {
		log.Printf("No results!")
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

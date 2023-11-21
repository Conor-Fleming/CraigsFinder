package craigslist

import (
	"fmt"

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
		panic(err)
	}

	if len(results) == 0 {
		fmt.Printf("No results!")
	}

	for _, v := range results {
		posting, err := v.Posting()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s - $%d\n", posting.Title, posting.Price)
		fmt.Printf("--------\n")
		fmt.Printf("%s\n", posting.Body)
	}
}

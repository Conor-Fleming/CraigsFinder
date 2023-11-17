package main

import (
	"fmt"

	"github.com/ecnepsnai/craigslist"
)

func main() {

	// Perform the search
	postings, err := craigslist.Search("apa", "2 bed", craigslist.LocationParams{
		AreaID:         1, //need a comprehensive list of AreaIDs, they dont seem so easy to find
		Latitude:       37.7749,
		Longitude:      122.4194,
		SearchDistance: 30, //in miles
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(postings) == 0 {
		fmt.Println("No results")
		return
	}

	// Print the results
	example, err := postings[0].Posting()
	if err != nil {
		fmt.Println("there was an error when getting posting detals", err)
	}

	fmt.Printf("%s - $%d\n", example.Title, example.Price)
	fmt.Printf("--------\n")
	fmt.Printf("%s\n", example.Body)
}

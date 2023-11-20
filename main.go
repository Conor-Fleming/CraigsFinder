package main

import (
	"fmt"

	"github.com/Conor-Flemign/CraigsFinder/craigslist"
)

func main() {
	// Populate site data
	allStates, err := craigslist.PopulateSiteData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display available states
	chosenState, err := craigslist.GetUserStateChoice(allStates)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Chosen State:", chosenState.RegName)
	// Use chosenState for further operations
}

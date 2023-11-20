package craigslist

import (
	"fmt"

	"github.com/sosedoff/go-craigslist"
)

type State struct {
	RegID   string
	RegName string
	Areas   []Area
}

type Area struct {
	ID   string
	Name string
	URL  string
}

func PopulateSiteData() (map[string]State, error) {
	sites, err := craigslist.Sites()
	if err != nil {
		return nil, err
	}

	states := make(map[string]State)

	for _, item := range sites {
		area := Area{
			ID:   item.Id,
			Name: item.Name,
			URL:  item.URL,
		}
		stateID := item.RegionId
		stateName := item.RegionName

		state, exists := states[stateID]
		if !exists {
			state = State{
				RegID:   stateID,
				RegName: stateName,
			}
		}

		state.Areas = append(state.Areas, area)
		states[stateID] = state
	}

	return states, nil
}

func DisplayStates(states map[string]State) {
	fmt.Println("Available States:")
	for _, state := range states {
		fmt.Printf("%s - %s\n", state.RegID, state.RegName)
	}
}

func GetUserStateChoice(states map[string]State) (*State, error) {
	DisplayStates(states)

	var userInput string
	fmt.Print("Enter the state ID you want to choose: ")
	fmt.Scanln(&userInput)

	chosenState, exists := states[userInput]
	if !exists {
		return nil, fmt.Errorf("state not found")
	}

	return &chosenState, nil
}

func GetUserAreaChoice(area *Area) {
	//might need to rethink this
}

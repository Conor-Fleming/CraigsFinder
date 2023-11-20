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

func PopulateSiteData() error {
	sites, err := craigslist.Sites()
	if err != nil {
		return err
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

		//if the state hasnt been added to the map, create a new State object
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

	fmt.Println(states["california"])
	return nil
}

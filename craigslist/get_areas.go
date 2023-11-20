package craigslist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SubArea struct {
	Abbreviation     string `json:"Abbreviation"`
	Description      string `json:"Description"`
	ShortDescription string `json:"ShortDescription"`
	SubAreaID        int    `json:"SubAreaID"`
}

type Area struct {
	Abbreviation     string    `json:"Abbreviation"`
	AreaID           int       `json:"AreaID"`
	Country          string    `json:"Country"`
	Description      string    `json:"Description"`
	Hostname         string    `json:"Hostname"`
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	Region           string    `json:"Region"`
	ShortDescription string    `json:"ShortDescription"`
	SubAreas         []SubArea `json:"SubAreas"`
	Timezone         string    `json:"Timezone"`
}

func FetchAreas(url string) ([]Area, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	var areas []Area
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	return areas, nil
}

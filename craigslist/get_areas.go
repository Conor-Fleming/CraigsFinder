package craigslist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

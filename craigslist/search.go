package craigslist

import "github.com/sosedoff/go-craigslist"

type Search struct {
	location []string
	options  craigslist.SearchOptions
}

func (s *Search) NewSearch(location []string, options craigslist.SearchOptions) Search {
	return Search{
		location: location,
		options:  options,
	}
}

func RunSearch()

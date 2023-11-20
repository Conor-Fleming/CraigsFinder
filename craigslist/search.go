package craigslist

import "github.com/sosedoff/go-craigslist"

type Search struct {
	location []string
	options  craigslist.SearchOptions
}

func (s *Search) NewSearch() Search {

	return Search{}
}

package craigslist

type SearchOptions struct {
	Category         string `form:"category"`
	Query            string `form:"query"`
	TitlesOnly       bool   `form:"titles_only"`
	HasImage         bool   `form:"has_image"`
	PostedToday      bool   `form:"posted_today"`
	BundleDuplicates bool   `form:"bundle_duplicates"`
	IncludeNearby    bool   `form:"include_nearby"`
	MinPrice         int    `form:"min_price"`
	MaxPrice         int    `form:"max_price"`
	Skip             int    `form:"skip"`
	Params           map[string]string
}

func (opts SearchOptions) CreateSearchOptions() {

}

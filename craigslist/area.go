package craigslist

type Area struct {
	Abbreviation     string    `json:"Abbreviation"`
	AreaID           int       `json:"AreaID"`
	Country          string    `json:"Country"`
	Description      string    `json:"Description"`
	Hostname         string    `json:"Hostname"`
	Latitude         float32   `json:"Latitude"`
	Longitude        float32   `json:"Longitude"`
	Region           string    `json:"Region"`
	ShortDescription string    `json:"ShortDescription"`
	SubAreas         []SubArea `json:"SubAreas"`
	Timezone         string    `json:"Timezone"`
}

type SubArea struct {
	Abbreviation     string `json:"Abbreviation"`
	Description      string `json:"Description"`
	ShortDescription string `json:"ShortDescription"`
	SubAreaID        int    `json:"SubAreaID"`
}

package hotel

import "time"

type OccupancyQuery struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type FindHotelsRatesQuery struct {
	CheckIn          time.Time
	CheckOut         time.Time
	Language         string
	Currency         string
	GuestNationality string
	HotelIds         []string
	Occupancies      []OccupancyQuery
}

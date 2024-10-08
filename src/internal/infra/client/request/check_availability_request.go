package request

type CheckAvailabilityRequest struct {
	Stay        Stay        `json:"stay"`
	Occupancies []Occupancy `json:"occupancies"`
	Hotels      Hotels      `json:"hotels"`
	Language    string      `json:"language"`
}

type Stay struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type Occupancy struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type Hotels struct {
	Hotel []int `json:"hotel"`
}

package response

type CheckAvailabilityResponse struct {
	Supplier       SupplierResponse `json:"supplier"`
	HotelsResponse HotelsWrapper    `json:"hotels"`
}

type HotelsWrapper struct {
	Hotels []HotelResponse `json:"hotels"`
	Total  int             `json:"total"`
}

type HotelResponse struct {
	Code     int            `json:"code"`
	Rooms    []RoomResponse `json:"rooms"`
	MinRate  string         `json:"minRate"`
	Currency string         `json:"currency"`
}

type RoomResponse struct {
	Name  string          `json:"name"`
	Rates []RatesResponse `json:"rates"`
}

type RatesResponse struct {
	RateKey string `json:"rateKey"`
	Net     string `json:"net"`
}

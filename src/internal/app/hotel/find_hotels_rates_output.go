package hotel

type FindHotelsRatesOutput struct {
	Data     []RateOutput   `json:"data"`
	Supplier SupplierOutput `json:"supplier"`
}

type RateOutput struct {
	HotelId  int     `json:"hotelId"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

type SupplierOutput struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

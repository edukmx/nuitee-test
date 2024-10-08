package request

import (
	"errors"
	"strings"
	"time"
)

type OccupancyRequest struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type HotelRatesRequest struct {
	CheckIn          time.Time `form:"checkin" binding:"required" time_format:"2006-01-02"`
	CheckOut         time.Time `form:"checkout" binding:"required" time_format:"2006-01-02"`
	Currency         string    `form:"currency" binding:"required"`
	GuestNationality string    `form:"guestNationality" binding:"required"`
	HotelIds         []string
	Occupancies      []OccupancyRequest
}

func (r *HotelRatesRequest) Validate() error {
	var validationErrors []string
	currentTime := time.Now()

	// Check if CheckIn date is in the future
	if r.CheckIn.Before(currentTime) {
		validationErrors = append(validationErrors, "checkIn date must be in the future")
	}

	// Check if CheckOut date is in the future
	if r.CheckOut.Before(currentTime) {
		validationErrors = append(validationErrors, "checkOut date must be in the future")
	}

	// Check if CheckOut is after CheckIn
	if r.CheckOut.Before(r.CheckIn) {
		validationErrors = append(validationErrors, "checkOut date must be after CheckIn date")
	}

	if r.Currency != "EUR" {
		validationErrors = append(validationErrors, "currency must be EUR")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

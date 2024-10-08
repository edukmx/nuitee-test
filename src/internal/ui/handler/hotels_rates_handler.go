package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/edukmx/nuitee/internal/app"
	"github.com/edukmx/nuitee/internal/app/hotel"
	"github.com/edukmx/nuitee/internal/domain"
	"github.com/edukmx/nuitee/internal/ui/request"
	"github.com/gin-gonic/gin"
)

type HotelsRatesHandler struct {
	findRatesHandler app.QueryHandler[hotel.FindHotelsRatesQuery, hotel.FindHotelsRatesOutput]
}

func NewHotelsRatesHandler(
	findRatesHandler app.QueryHandler[hotel.FindHotelsRatesQuery, hotel.FindHotelsRatesOutput],
) *HotelsRatesHandler {
	return &HotelsRatesHandler{
		findRatesHandler: findRatesHandler,
	}
}

func (h *HotelsRatesHandler) List(c *gin.Context) {
	var hotelRatesRequest request.HotelRatesRequest

	// Mapping query params to the request object
	if err := c.ShouldBindQuery(&hotelRatesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	occupancies := c.Query("occupancies")
	if occupancies == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "occupancies is required"})
		return
	}

	if err := json.Unmarshal([]byte(occupancies), &hotelRatesRequest.Occupancies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid occupancies format"})
		return
	}

	hotelIds := c.Query("hotelIds")
	if hotelIds != "" {
		hotelRatesRequest.HotelIds = strings.Split(hotelIds, ",")
	}

	if err := hotelRatesRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var occupanciesQuery []hotel.OccupancyQuery

	query := h.queryFromRequest(hotelRatesRequest, occupanciesQuery)

	res, err := h.findRatesHandler.Ask(
		query,
	)
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(domain.ErrNationalityNotSupported, err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Nationality"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *HotelsRatesHandler) queryFromRequest(
	hotelRatesRequest request.HotelRatesRequest,
	occupanciesQuery []hotel.OccupancyQuery,
) hotel.FindHotelsRatesQuery {
	for _, occupancy := range hotelRatesRequest.Occupancies {
		occupanciesQuery = append(occupanciesQuery, hotel.OccupancyQuery{
			Adults:   occupancy.Adults,
			Rooms:    occupancy.Rooms,
			Children: occupancy.Children,
		})
	}
	return hotel.FindHotelsRatesQuery{
		CheckIn:          hotelRatesRequest.CheckIn,
		CheckOut:         hotelRatesRequest.CheckOut,
		Currency:         hotelRatesRequest.Currency,
		Occupancies:      occupanciesQuery,
		HotelIds:         hotelRatesRequest.HotelIds,
		GuestNationality: hotelRatesRequest.GuestNationality,
	}
}

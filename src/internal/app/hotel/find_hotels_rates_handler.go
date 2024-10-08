package hotel

import (
	"strconv"

	"github.com/edukmx/nuitee/internal/app"
	"github.com/edukmx/nuitee/internal/domain/nationality/service"
	"github.com/edukmx/nuitee/internal/infra/client"
	"github.com/edukmx/nuitee/internal/infra/client/request"
)

type FindHotelsRatesHandler struct {
	nationalityFinder *service.FindByCode
	apiClient         client.HotelBedsClient
}

func NewFindHotelsRatesHandler(
	nationalityFinder *service.FindByCode,
	apiClient client.HotelBedsClient,
) app.QueryHandler[FindHotelsRatesQuery, FindHotelsRatesOutput] {
	return &FindHotelsRatesHandler{
		nationalityFinder: nationalityFinder,
		apiClient:         apiClient,
	}
}

func (h *FindHotelsRatesHandler) Ask(query FindHotelsRatesQuery) (*FindHotelsRatesOutput, error) {

	nationalityObj, err := h.nationalityFinder.Find(query.GuestNationality)
	if err != nil {
		return nil, err
	}

	resp, err := h.apiClient.CheckAvailability(h.transformQueryToCheckAvailabilityRequest(query, nationalityObj.LanguageISO))
	if err != nil {
		return nil, err
	}

	var data []RateOutput
	for _, hotel := range resp.HotelsResponse.Hotels {
		price, err := strconv.ParseFloat(hotel.MinRate, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, RateOutput{
			HotelId:  hotel.Code,
			Currency: hotel.Currency,
			Price:    price,
		})
	}

	response := FindHotelsRatesOutput{
		Data: data,
		Supplier: SupplierOutput{
			Request:  resp.Supplier.Request,
			Response: resp.Supplier.Response,
		},
	}

	return &response, nil
}

func (h *FindHotelsRatesHandler) transformQueryToCheckAvailabilityRequest(
	query FindHotelsRatesQuery,
	lang string,
) *request.CheckAvailabilityRequest {

	stay := request.Stay{
		CheckIn:  query.CheckIn.Format("2006-01-02"),
		CheckOut: query.CheckOut.Format("2006-01-02"),
	}

	var hotelIds []int
	for _, id := range query.HotelIds {
		if hotelId, err := strconv.Atoi(id); err == nil {
			hotelIds = append(hotelIds, hotelId)
		}
	}

	var occupancies []request.Occupancy
	for _, occupancy := range query.Occupancies {
		occupancies = append(occupancies, request.Occupancy{
			Rooms:    occupancy.Rooms,
			Adults:   occupancy.Adults,
			Children: occupancy.Children,
		})
	}

	// omitting language, the format is not known, iso code not accepted
	return &request.CheckAvailabilityRequest{
		Stay:        stay,
		Occupancies: occupancies,
		Hotels:      request.Hotels{Hotel: hotelIds},
		// Language: lang,
	}
}

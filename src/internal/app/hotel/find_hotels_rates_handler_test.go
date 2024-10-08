package hotel_test

import (
	"errors"
	"testing"
	"time"

	"github.com/edukmx/nuitee/internal/app/hotel"
	"github.com/edukmx/nuitee/internal/domain/nationality"
	"github.com/edukmx/nuitee/internal/domain/nationality/service"
	"github.com/edukmx/nuitee/internal/infra/client/response"
	"github.com/edukmx/nuitee/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindHotelsRatesHandler_Ask(t *testing.T) {
	// Define multiple test cases
	testCases := []struct {
		name             string
		query            hotel.FindHotelsRatesQuery
		nationality      *nationality.Nationality
		nationalityError error
		apiResponse      *response.CheckAvailabilityResponse
		apiError         error
		expectedOutput   *hotel.FindHotelsRatesOutput
		expectedError    error
	}{
		{
			name: "it should return output with no errors",
			query: hotel.FindHotelsRatesQuery{
				GuestNationality: "AR",
				CheckIn:          time.Now(),
				CheckOut:         time.Now().Add(24 * time.Hour),
				HotelIds:         []string{"1001"},
				Occupancies: []hotel.OccupancyQuery{
					{Rooms: 1, Adults: 2, Children: 0},
				},
			},
			nationality: &nationality.Nationality{
				NationalityISO: "AR",
				LanguageISO:    "es",
			},
			nationalityError: nil,
			apiResponse: &response.CheckAvailabilityResponse{
				HotelsResponse: response.HotelsWrapper{
					Hotels: []response.HotelResponse{
						{
							Code:     1001,
							MinRate:  "200.00",
							Currency: "USD",
						},
					},
				},
				Supplier: response.SupplierResponse{
					Request:  "request_data",
					Response: "response_data",
				},
			},
			apiError: nil,
			expectedOutput: &hotel.FindHotelsRatesOutput{
				Data: []hotel.RateOutput{
					{
						HotelId:  1001,
						Currency: "USD",
						Price:    200.00,
					},
				},
				Supplier: hotel.SupplierOutput{
					Request:  "request_data",
					Response: "response_data",
				},
			},
			expectedError: nil,
		},
		{
			name: "it should return not found error, when nationality is wrong",
			query: hotel.FindHotelsRatesQuery{
				GuestNationality: "XX",
			},
			nationality:      nil,
			nationalityError: errors.New("nationality not found"),
			apiResponse:      nil,
			apiError:         nil,
			expectedOutput:   nil,
			expectedError:    errors.New("nationality not found"),
		},
		{
			name: "it should return not found error, when external api fails",
			query: hotel.FindHotelsRatesQuery{
				GuestNationality: "AR",
			},
			nationality: &nationality.Nationality{
				NationalityISO: "AR",
				LanguageISO:    "es",
			},
			nationalityError: nil,
			apiResponse:      nil,
			apiError:         errors.New("API error"),
			expectedOutput:   nil,
			expectedError:    errors.New("API error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockNationalityRepository := new(mocks.Repository)
			mockApiClient := new(mocks.HotelBedsClient)

			mockNationalityFinder := service.NewFindByCode(mockNationalityRepository)
			handler := hotel.NewFindHotelsRatesHandler(mockNationalityFinder, mockApiClient)

			mockNationalityRepository.
				On("FindByIso", tc.query.GuestNationality).
				Return(tc.nationality, tc.nationalityError)

			if tc.nationality != nil {
				mockApiClient.On("CheckAvailability", mock.Anything).Return(tc.apiResponse, tc.apiError)
			}

			output, err := handler.Ask(tc.query)

			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, output)
			}

			mockNationalityRepository.AssertExpectations(t)
			mockApiClient.AssertExpectations(t)
		})

	}
}

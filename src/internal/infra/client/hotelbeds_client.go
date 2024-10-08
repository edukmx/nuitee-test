package client

import (
	"github.com/edukmx/nuitee/internal/infra/client/request"
	"github.com/edukmx/nuitee/internal/infra/client/response"
)

//go:generate mockery --case=snake --outpkg=mocks --output=../../mocks --name=HotelBedsClient
type HotelBedsClient interface {
	CheckAvailability(reqObj *request.CheckAvailabilityRequest) (*response.CheckAvailabilityResponse, error)
}

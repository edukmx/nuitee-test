package client_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/edukmx/nuitee/config"
	"github.com/edukmx/nuitee/internal/infra/client"
	"github.com/edukmx/nuitee/internal/infra/client/request"
	"github.com/edukmx/nuitee/internal/infra/client/response"
	"github.com/stretchr/testify/assert"
)

func TestHotelBedsAdapter_CheckAvailability_Success(t *testing.T) {

	mockConfig := &config.Config{
		HotelBedsConfig: config.HotelBedsConfig{
			Host:   "http://test-host",
			ApiKey: "test-api-key",
			Secret: "test-secret",
		},
		ClientTimeout: 2 * time.Second,
	}

	reqObj := &request.CheckAvailabilityRequest{}

	expectedResponse := &response.CheckAvailabilityResponse{
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
			Request:  "test-request",
			Response: "test-response",
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "test-api-key", r.Header.Get("Api-key"))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer mockServer.Close()

	mockConfig.HotelBedsConfig.Host = mockServer.URL

	adapter := client.NewHotelBedsAdapter(mockConfig)

	actualResponse, err := adapter.CheckAvailability(reqObj)

	assert.NoError(t, err)

	expectedRequest, _ := json.Marshal(reqObj)
	expectedResponseBody, _ := json.Marshal(expectedResponse)

	assert.JSONEq(t, string(expectedRequest), actualResponse.Supplier.Request)
	assert.JSONEq(t, string(expectedResponseBody), actualResponse.Supplier.Response)

	assert.Equal(t, expectedResponse.HotelsResponse, actualResponse.HotelsResponse)
}

func TestHotelBedsAdapter_CheckAvailability_Failure(t *testing.T) {

	mockConfig := &config.Config{
		HotelBedsConfig: config.HotelBedsConfig{
			Host:   "http://test-host",
			ApiKey: "test-api-key",
			Secret: "test-secret",
		},
		ClientTimeout: 2 * time.Second,
	}

	reqObj := &request.CheckAvailabilityRequest{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"something went wrong"}`)
	}))
	defer mockServer.Close()

	mockConfig.HotelBedsConfig.Host = mockServer.URL

	adapter := client.NewHotelBedsAdapter(mockConfig)

	actualResponse, err := adapter.CheckAvailability(reqObj)

	assert.Error(t, err)
	assert.Nil(t, actualResponse)
}

func TestHotelBedsAdapter_GenerateSignature(t *testing.T) {

	mockConfig := &config.Config{
		HotelBedsConfig: config.HotelBedsConfig{
			ApiKey: "test-api-key",
			Secret: "test-secret",
		},
	}

	adapter := client.NewHotelBedsAdapter(mockConfig)

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	data := mockConfig.HotelBedsConfig.ApiKey + mockConfig.HotelBedsConfig.Secret + timestamp
	expectedHash := sha256.Sum256([]byte(data))
	expectedSignature := hex.EncodeToString(expectedHash[:])

	actualSignature := adapter.(*client.HotelBedsAdapter).GenerateSignature(
		mockConfig.HotelBedsConfig.ApiKey,
		mockConfig.HotelBedsConfig.Secret,
	)

	assert.Equal(t, expectedSignature, actualSignature)
}

func TestResponseMapping(t *testing.T) {
	file, err := getSuccessResponse()
	if err != nil {
		t.Error(err)
	}
	var output response.CheckAvailabilityResponse
	err = json.Unmarshal(file, &output)
	if err != nil {
		t.Error(err)
	}
}

func getSuccessResponse() ([]byte, error) {
	data, err := os.ReadFile("../../../resources/availability_response.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

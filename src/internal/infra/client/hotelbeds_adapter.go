package client

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/edukmx/nuitee/config"
	"github.com/edukmx/nuitee/internal/infra/client/request"
	"github.com/edukmx/nuitee/internal/infra/client/response"
)

type HotelBedsAdapter struct {
	config *config.Config
	client *http.Client
}

func NewHotelBedsAdapter(config *config.Config) HotelBedsClient {
	return &HotelBedsAdapter{
		config: config,
		client: &http.Client{
			Timeout: config.ClientTimeout,
		},
	}
}

func (h HotelBedsAdapter) CheckAvailability(reqObj *request.CheckAvailabilityRequest) (
	*response.CheckAvailabilityResponse,
	error,
) {
	url := fmt.Sprintf("%s/hotel-api/1.0/hotels", h.config.HotelBedsConfig.Host)

	reqBody, err := json.Marshal(reqObj)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", h.GenerateSignature(h.config.HotelBedsConfig.ApiKey, h.config.HotelBedsConfig.Secret))
	req.Header.Set("Api-key", h.config.HotelBedsConfig.ApiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error(fmt.Sprintf("error in request: %v", string(respBody)))
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var output response.CheckAvailabilityResponse

	if err := json.Unmarshal(respBody, &output); err != nil {
		return nil, fmt.Errorf("no rates found for the given request")
	}

	output.Supplier.Response = string(respBody)
	output.Supplier.Request = string(reqBody)
	return &output, err

}

func (h HotelBedsAdapter) GenerateSignature(apiKey, secret string) string {
	// Get the current timestamp in seconds
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	// Concatenate the apiKey, secret, and timestamp
	data := apiKey + secret + timestamp

	// Generate the SHA-256 hash
	hash := sha256.Sum256([]byte(data))

	// Convert the hash to a hexadecimal string
	signature := hex.EncodeToString(hash[:])

	return signature
}

package geoLocationProvider

import (
	"GeO-Locator/internal/customErrors"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

// Mock Http Transport
type mockHttpTransport struct {
	mockResponse *http.Response
	mockError    error
}

func (m *mockHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.mockResponse, m.mockError
}

func TestIpInfoGeoLocationProvider_GetGeoLocation_WhenIpNotExists_PanicShouldBeThrown(t *testing.T) {
	defer func() {
		expectedError := customErrors.NewCodedCustomError("GeoLocation not found for '1.1.1.1'", 404)

		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				var codedErr *customErrors.CodedCustomError
				if errors.As(err, &codedErr) && codedErr.Error() == expectedError.Error() {
					return
				}
			}

			t.Errorf("Expected CodedCustomError{Message: 'GeoLocation not found for '1.1.1.1'', Code: 404}, got %v", r)
		}
	}()

	// mock settings
	settings := ipInfoSettings{
		ApiKey: "fake-api-key",
		ApiUrl: "http://fakeapi.com",
	}

	// mock HTTP response
	mockResponse := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
	}

	// mock HTTP client
	mockClient := &http.Client{
		Transport: &mockHttpTransport{mockResponse: mockResponse},
	}

	// sut object
	provider := ipInfoGeoLocationProvider{
		httpClient: mockClient,
		settings:   settings,
	}

	_ = provider.GetGeoLocation("1.1.1.1")
}

func TestIpInfoGeoLocationProvider_GetGeoLocation_WhenIpExists_ReturnsGeoLocation(t *testing.T) {
	// mock settings
	settings := ipInfoSettings{
		ApiKey: "fake-api-key",
		ApiUrl: "http://fakeapi.com",
	}

	// mock HTTP response object
	expectedGeoLocation := GeoLocation{
		City:    "Istanbul",
		Country: "Turkey",
	}

	responseBody, _ := json.Marshal(expectedGeoLocation)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(responseBody)),
	}

	// mock HTTP client
	mockClient := &http.Client{
		Transport: &mockHttpTransport{mockResponse: mockResponse},
	}

	provider := ipInfoGeoLocationProvider{
		httpClient: mockClient,
		settings:   settings,
	}

	geoLocation := provider.GetGeoLocation("1.1.1.1")

	if geoLocation.City != expectedGeoLocation.City || geoLocation.Country != expectedGeoLocation.Country {
		t.Errorf("Expected %v, got %v", expectedGeoLocation, geoLocation)
	}
}

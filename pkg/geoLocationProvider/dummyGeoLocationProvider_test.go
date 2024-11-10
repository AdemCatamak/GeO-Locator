package geoLocationProvider

import (
	"GeO-Locator/internal/customErrors"
	"errors"
	"testing"
)

func TestDummyGeoLocationProvider_WhenIpExists_ResponseShouldNotBeNil(t *testing.T) {

	data := map[string]*GeoLocation{
		"1.1.1.1": {
			Country: "Country1",
			City:    "City1",
		},
	}
	provider := dummyGeoLocationProvider{
		data: data,
	}

	geoLocation := provider.GetGeoLocation("1.1.1.1")
	if geoLocation.Country != "Country1" || geoLocation.City != "City1" {
		t.Errorf("Expected GeoLocation{Country: 'Country1', City: 'City1'}, got %v", geoLocation)
	}
}

func TestDummyGeoLocationProvider_WhenIpDoesNotExist_ErrorShouldBeThrown(t *testing.T) {
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

	provider := dummyGeoLocationProvider{
		data: map[string]*GeoLocation{},
	}

	provider.GetGeoLocation("1.1.1.1")
}

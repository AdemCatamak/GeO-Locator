package geoLocationProvider

import (
	"GeO-Locator/internal/config"
	"GeO-Locator/internal/customErrors"
	"fmt"
)

type dummyGeoLocationProvider struct {
	data map[string]*GeoLocation
}

func newDummyGeoLocationProvider() dummyGeoLocationProvider {
	configManager := config.GetConfigManager()
	var locationList []GeoLocationCollectionItem
	configManager.GetObj("Locations", &locationList)

	data := map[string]*GeoLocation{}
	for _, item := range locationList {
		data[item.Ip] = &item.GeoLocation
	}

	return dummyGeoLocationProvider{
		data: data,
	}
}

func (d dummyGeoLocationProvider) GetGeoLocation(ip string) *GeoLocation {
	geoLocation := d.data[ip]
	if geoLocation == nil {
		panic(customErrors.NewCodedCustomError(fmt.Sprintf("GeoLocation not found for '%s'", ip), 404))
	}

	return geoLocation
}

type GeoLocationCollectionItem struct {
	Ip          string      `json:"Ip"`
	GeoLocation GeoLocation `json:"GeoLocation"`
}

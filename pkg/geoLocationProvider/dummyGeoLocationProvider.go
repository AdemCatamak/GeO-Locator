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
	configManager.GetObj(config.LocationsKey, &locationList)

	data := map[string]*GeoLocation{}
	for _, item := range locationList {
		data[item.Ip] = &item.GeoLocation
	}

	return dummyGeoLocationProvider{
		data: data,
	}
}

func (p dummyGeoLocationProvider) GetGeoLocation(ip string) *GeoLocation {
	geoLocation := p.data[ip]
	if geoLocation == nil {
		panic(customErrors.NewCodedCustomError(fmt.Sprintf("GeoLocation not found for '%s'", ip), 404))
	}

	return geoLocation
}

type GeoLocationCollectionItem struct {
	Ip          string      `json:"Ip"`
	GeoLocation GeoLocation `json:"GeoLocation"`
}

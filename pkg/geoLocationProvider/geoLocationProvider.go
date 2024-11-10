package geoLocationProvider

import "GeO-Locator/internal/config"

type GeoLocationProvider interface {
	GetGeoLocation(ip string) *GeoLocation
}

type GeoLocation struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func NewGeoLocationProvider() GeoLocationProvider {
	configManager := config.GetConfigManager()
	selectedProvider := configManager.GetInt(config.SelectedGeoLocatorProviderKey)

	switch selectedProvider {
	case ProviderDummy:
		return newDummyGeoLocationProvider()
	case ProviderIpInfo:
		return newIpInfoGeoLocationProvider()
	}

	panic("Unknown geolocation provider selection")
}

const (
	ProviderDummy = iota + 1
	ProviderIpInfo
)

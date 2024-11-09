package geoLocationProvider

type GeoLocationProvider interface {
	GetGeoLocation(ip string) *GeoLocation
}

type GeoLocation struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func NewGeoLocationProvider() GeoLocationProvider {
	return newDummyGeoLocationProvider()
}

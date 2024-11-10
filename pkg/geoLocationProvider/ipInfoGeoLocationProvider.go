package geoLocationProvider

import (
	"GeO-Locator/internal/config"
	"GeO-Locator/internal/customErrors"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type ipInfoGeoLocationProvider struct {
	httpClient *http.Client
	settings   ipInfoSettings
}

func newIpInfoGeoLocationProvider() ipInfoGeoLocationProvider {

	settings := ipInfoSettings{}

	configManager := config.GetConfigManager()
	configManager.GetObj(config.IpInfoSettingsKey, &settings)

	return ipInfoGeoLocationProvider{
		httpClient: &http.Client{},
		settings:   settings,
	}
}

func (p ipInfoGeoLocationProvider) GetGeoLocation(ip string) *GeoLocation {
	encodedIp := url.PathEscape(ip)
	httpReq, _ := http.NewRequest("GET", p.settings.ApiUrl+"/"+encodedIp+"/json", nil)

	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Authorization", "Bearer "+p.settings.ApiKey)

	httpRes, err := p.httpClient.Do(httpReq)
	if err != nil {
		panic(err)
	}

	if httpRes.StatusCode == http.StatusNotFound {
		panic(customErrors.NewCodedCustomError(fmt.Sprintf("GeoLocation not found for '%s'", ip), 404))
	} else if httpRes.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Unexpected status code: %d", httpRes.StatusCode))
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(httpRes.Body)

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Body")
	log.Println(string(body))

	var geoLocation GeoLocation
	err = json.Unmarshal(body, &geoLocation)
	if err != nil {
		panic(err)
	}

	return &geoLocation
}

type ipInfoSettings struct {
	ApiKey string `json:"ApiKey"`
	ApiUrl string `json:"ApiUrl"`
}

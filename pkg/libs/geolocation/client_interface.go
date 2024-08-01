package geolocation

import (
	"log"
	"net/http"
)

type GeoLocationClient interface {
	Url(ip string) string
	GetGeolocation(ip string) (*http.Response, error)
}

type IpBasedClient struct {
	url string
}

func (client IpBasedClient) Url(ip string) string {
	return client.url + "/" + ip
}

func (client IpBasedClient) GetGeolocation(ip string) (*http.Response, error) {
	fullURL := client.Url(ip)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Println("Cannot create request:", err)
		return nil, err
	}
	// Execute the request using an HTTP client
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("Cannot complete request:", err)
		return nil, err
	}
	return res, nil
}

// Package geolocation provides functionalities for geolocation lookups.
//
// This package contains types and functions for performing geolocation lookups
// based on IP addresses and returning geolocation information.
package geolocation

import (
	"encoding/json"
	"log"
)

type GeoLocation struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"countryCode"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
}

// Create a client
func NewIpBasedClient() GeoLocationClient {
	return IpBasedClient{url: "http://ip-api.com/json"}
}

// GetGeoLocation performs a geolocation lookup based on the provided IP address.
//
// Parameters:
// - ip: The IP address to look up.
//
// Returns:
// - GeoLocation: A struct containing the geolocation information for the IP address.
// - error: An error if the lookup fails.
func GetGeoLocationIpBased(ip string) (*GeoLocation, error) {
	client := NewIpBasedClient()
	res, err := client.GetGeolocation(ip)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// Parse the response body
	var geoLocationResponse GeoLocation
	if err := json.NewDecoder(res.Body).Decode(&geoLocationResponse); err != nil {
		log.Println("Cannot parse response:", err)
		return nil, err
	}
	// Manually add the IP to the response
	geoLocationResponse.IP = ip

	return &geoLocationResponse, nil
}

func twoSum(nums []int, target int) []int {
	memoization := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := memoization[complement]; found {
			return []int{index, i}
		}
		memoization[num] = i
	}
	return []int{}
}

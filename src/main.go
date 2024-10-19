package main

import (
	"sort"
)

func main() {
	var allIPs []string = readIPs()

	var allCities = getAllCities(allIPs)

	// get a slcie of all unique cities
	var citiesSlice []city
	for _, c := range allCities {
		citiesSlice = append(citiesSlice, c)
	}

	// sort by number of requests
	sort.SliceStable(citiesSlice, func(i, j int) bool {
		return citiesSlice[i].numResquests > citiesSlice[j].numResquests
	})
	outputNumberOfRequests(citiesSlice)

	// calculate the distance
	for i := 0; i < len(citiesSlice); i++ {
		citiesSlice[i].distance = citiesSlice[i].calculateDistanceFromVictoria()
	}

	// sort by distance
	sort.SliceStable(citiesSlice, func(i, j int) bool {
		return citiesSlice[i].distance > citiesSlice[j].distance
	})
	outputFurthestCity(citiesSlice[0])
}

func getAllCities(allIPs []string) map[string]city {

	var allCities = make(map[string]city)

	// for each ip
	for _, ip := range allIPs {
		// parse the attributes of the ip
		var cityName string = getCityName(ip)
		var country string = getCountry(ip)
		lat, long := getLocation(ip)

		// only interested in cities outside of canada
		if country == "Canada" {
			continue
		}

		c, exists := allCities[cityName]

		// update the map of cities
		if exists {
			newRequestNum := c.numResquests + 1
			var newCity city = city{c.Name, c.country, newRequestNum, 0, long, lat}
			allCities[cityName] = newCity
		} else {
			var newCity city = city{cityName, country, 1, 0, long, lat}
			allCities[cityName] = newCity
		}
	}

	return allCities
}

package main

import (
	"fmt"
	"sort"
	//"bufio"
)

func main() {
	var allIPs []string = readIPs()

	var allCities = getAllCities(allIPs)

	var citiesSlice []city
	for _, c := range allCities {
		citiesSlice = append(citiesSlice, c)
	}
	
	//sort
	sort.SliceStable(citiesSlice, func(i, j int) bool {
		return citiesSlice[i].numResquests > citiesSlice[j].numResquests
	})

	fmt.Printf("slice of cities: %v", citiesSlice)
	fmt.Println("---------------------------------------------------------------")


	for i:=0; i<len(citiesSlice); i++ {
		citiesSlice[i].distance = citiesSlice[i].calculateDistanceFromVictoria()
	}
	//for distnace, need to calculate
	sort.SliceStable(citiesSlice, func(i, j int) bool {
		return citiesSlice[i].distance > citiesSlice[j].distance
	})
	fmt.Printf("Sort by distance: %v", citiesSlice)
	// outputNumberOfRequests(citiesSlice)
	//put into other files
	// file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
	// datawriter := bufio.NewWriter(file)
 
	// for _, data := range sampledata {
	// 	_, _ = datawriter.WriteString(data + "\n")
	// }
 
	// datawriter.Flush()
	// file.Close()
	
}



func getAllCities(allIPs []string) map[string]city{

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

	return allCities;
}


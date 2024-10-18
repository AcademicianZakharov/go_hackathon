package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"strconv"
	"github.com/ipinfo/go/v2/ipinfo"
)

func readIPs() []string {
	var allIP = []string{}

	// read the log
	file, err := os.Open("ca-test-access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// for each line
	for scanner.Scan() {
		var ip string = strings.Split(scanner.Text(), " - - ")[0]
		allIP = append(allIP, ip)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allIP
}

func getCityName(ip string) string {
	const token = "177a20b5bbc6a6"

	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip))

	if err != nil {
		log.Fatal(err)
	}

	return info.City
}

func getLocation(ip string) (latitude float64, longitude float64) {
	const token = "177a20b5bbc6a6"

	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip))

	if err != nil {
		log.Fatal(err)
	}

	location := info.Location

	latitudeStr := strings.Split(location, ",")[0]
	longitudeStr := strings.Split(location, ",")[1]
	

	value1, err1 := strconv.ParseFloat(latitudeStr, 64)
	if err1 != nil {
		// do something sensible
	}
	latitude = float64(value1)

	value2, err2 := strconv.ParseFloat(longitudeStr, 64)
	if err2 != nil {
		// do something sensible
	}
	longitude = float64(value2)


	return
}

func getCountry(ip string) string {
	const token = "177a20b5bbc6a6"

	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip))

	if err != nil {
		log.Fatal(err)
	}

	return info.Country
}

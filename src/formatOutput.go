package main

import (
	"fmt"
	"os"
	"strconv"
)

var resultFile string = "../output/result.txt"

func outputNumberOfRequests(citiesSlice []city) {
	f, err := os.Create(resultFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	f.WriteString("Top 10 cities outside of Canada where the requests are from: \n")

	minCount := min(len(citiesSlice), 10)
	for i := 0; i < minCount; i++ {
		f.WriteString(citiesSlice[i].Name)
		f.WriteString("  |  ")
		f.WriteString(strconv.Itoa(citiesSlice[i].numResquests))
		f.WriteString(" requests")
		f.WriteString("\n")
	}

	f.WriteString("\n ------------------------------------\n\n")

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func outputFurthestCity(furthestCity city) {
	f, err := os.OpenFile(resultFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		f, err = os.Create(resultFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	f.WriteString("Of the top 10 cities, the city that is the furthest away from Victoria, BC is: ")

	f.WriteString(furthestCity.Name)

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

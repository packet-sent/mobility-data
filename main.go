package main

import (
	"fmt"

	"github.com/packet-sent/mobility-data/geospatial"
	"github.com/packet-sent/mobility-data/graphql"
)

func main() {
	/*
		Data can be requested via premade GraphQL queries
		Use the simpler queries when possible to speed up the requests and spend less API credits

		Available Queries:
		- simpleVehicle (vehicle type)
		- advancedVehicle (vehicle details)
		- simpleProvider (provider name)
		- advancedProvider (provider details)
	*/

	queryMode := "simpleVehicle"
	searchLocation := "WC2R"

	//GeoPostcode("WC2R 2PH")
	lat, lng, err := geospatial.GeoDistrict("WC2R")
	if err != nil {
		panic(err)
	}
	fmt.Println("Found geospatial data for", searchLocation)

	fmt.Println("Sending GraphQL data request for", searchLocation)
	graphql.RequestData(lat, lng, queryMode)

}

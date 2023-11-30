package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/machinebox/graphql"
	"github.com/tidwall/gjson"
)

func RequestData(lat, lng, mode, token string) interface{} {
	var graphQuery string
	var graphResponse interface{}

	switch mode {
	case "simpleVehicle":
		graphQuery = SimpleVehicleRequest
		graphResponse = SimpleVehicleResponse{}
	case "advancedVehicle":
		graphQuery = AdvancedVehicleRequest
		graphResponse = AdvancedVehicleResponse{}
	case "simpleProvider":
		graphQuery = SimpleProviderRequest
		graphResponse = SimpleVehicleResponse{}
	case "advancedProvider":
		graphQuery = AdvancedProviderRequest
		graphResponse = AdvancedProviderResponse{}
	default:
		log.Fatalln("Unknown mode, please use an available mode")
	}
	//Adds the search filters to the query
	graphQuery = strings.Replace(graphQuery, "$lat", lat, -1)
	graphQuery = strings.Replace(graphQuery, "$lng", lng, -1)

	client := graphql.NewClient("https://flow-api.fluctuo.com/v1?access_token=" + token)

	// ONLY USED FOR DEBUG PURPOSES
	//client.Log = func(s string) { log.Println(s) }

	req := graphql.NewRequest(graphQuery)

	//Avoid cached results
	req.Header.Set("Cache-Control", "no-cache")

	if err := client.Run(context.Background(), req, &graphResponse); err != nil {
		log.Fatal(err)
	}

	prettyJSON, err := json.MarshalIndent(graphResponse, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Full Data", string(prettyJSON))

	if mode == "simpleVehicle" || mode == "advancedVehicle" {
		data := gjson.Get(string(prettyJSON), "vehicles").Array()

		fmt.Println("Found", len(data), "shared mobility vehicles")
	} else {
		data := gjson.Get(string(prettyJSON), "providers").Array()

		fmt.Println("Found", len(data), "shared mobility providers")
	}
	return graphResponse
}

package graphql

var SimpleVehicleRequest = `
query Vehicles {
	vehicles(lat: $lat, lng: $lng) {
		id
		type
	}
}
`

var AdvancedVehicleRequest = `
query Vehicles {
    vehicles(lat: $lat, lng: $lng) {
        id
        lat
        lng
        type
        attributes
        propulsion
        battery
        pricing {
            currency
            unlock
        }
    }
}
`

var SimpleProviderRequest = `
query Providers {
    providers(lat: $lat, lng: $lng) {
        name
        website
    }
}
`

var AdvancedProviderRequest = `
query Providers {
    providers(lat: $lat, lng: $lng) {
        name
        slug
        website
        discountCode
        vehicleTypes
        freefloatingVehiclesTypes
        stationVehicleTypes
        carClasses
        enabled
    }
}
`

type SimpleVehicleResponse struct {
	Vehicles []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"vehicles"`
	Providers []struct{} `json:"omitempty"`
}

type AdvancedVehicleResponse struct {
	Vehicles []struct {
		ID         string   `json:"id"`
		Lat        float64  `json:"lat"`
		Lng        float64  `json:"lng"`
		Type       string   `json:"type"`
		Attributes []string `json:"attributes"`
		Propulsion string   `json:"propulsion"`
		Battery    int      `json:"battery"`
		Pricing    struct {
			Currency string  `json:"currency"`
			Unlock   float64 `json:"unlock"`
		} `json:"pricing"`
	} `json:"vehicles"`
	Providers []struct{} `json:"omitempty"`
}

type SimpleProviderResponse struct {
	Providers []struct {
		Name    string `json:"name"`
		Website string `json:"website"`
	} `json:"providers"`
	Vehicles []struct{} `json:"omitempty"`
}

type AdvancedProviderResponse struct {
	Providers []struct {
		Name                      string   `json:"name"`
		Slug                      string   `json:"slug"`
		Website                   string   `json:"website"`
		VehicleTypes              []string `json:"vehicleTypes"`
		FreefloatingVehiclesTypes []string `json:"freefloatingVehiclesTypes"`
		StationVehicleTypes       []string `json:"stationVehicleTypes"`
		Enabled                   bool     `json:"enabled"`
	} `json:"providers"`
	Vehicles []struct{} `json:"omitempty"`
}

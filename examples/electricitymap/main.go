package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/thegreenwebfoundation/grid-intensity-go/pkg/provider"
)

func main() {
	// Register at https://api-portal.electricitymaps.com/
	token := os.Getenv("ELECTRICITY_MAPS_API_TOKEN")
	if token == "" {
		log.Fatalln("please set the env variable `ELECTRICITY_MAPS_API_TOKEN`")
	}
	url := os.Getenv("ELECTRICITY_MAPS_API_URL")
	if url != "" {
		log.Fatalln("please set the env variable `ELECTRICITY_MAPS_API_URL`")
	}

	c := provider.ElectricityMapsConfig{
		APIURL: url,
		Token:  token,
	}
	e, err := provider.NewElectricityMaps(c)
	if err != nil {
		log.Fatalln("could not make provider", err)
	}

	res, err := e.GetCarbonIntensity(context.Background(), "AU-SA")
	if err != nil {
		log.Fatalln("could not get carbon intensity", err)
	}

	bytes, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		log.Fatalln("could not get carbon intensity", err)
	}

	fmt.Println(string(bytes))
}

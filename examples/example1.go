package main

import (
	"fmt"
	"github.com/mrsln/geocoder"
)

func main() {
	address := "Moscow, Kremlin"
	coords, err := geocoder.AddressToCoordinates(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Coordinates for '%s' are: %f (Latitude) %f (Longitude)\n",
		address, coords.Latitude, coords.Longitude)
}

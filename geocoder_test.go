package geocoder

import (
	"testing"
	"fmt"
)

func TestAddressToCoordinates(t *testing.T) {
	const address = "Moscow, Kremlin"
	coords, err := AddressToCoordinates(address)
	if err != nil {
		t.Fatal(err)
	}
	if coords.Latitude != 55.751999 || coords.Longitude != 37.617733 {
		t.Fatalf("Coordinates for %s are wrong: %f, %f", address, coords.Latitude, coords.Longitude)
	}
}

func Example_AddressToCoordinates() {
	const address = "Moscow, Kremlin"
	coords, _ := AddressToCoordinates(address)
	fmt.Printf("Coordinates for %s are: %f, %f\n", address, coords.Latitude, coords.Longitude)
	// Output: Coordinates for Moscow, Kremlin are: 55.751999, 37.617733
}

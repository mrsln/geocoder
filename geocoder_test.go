package geocoder

import "testing"

func TestAddressToCoordinates(t *testing.T) {
  const address = "Moscow, Kremlin"
  coords, err := AddressToCoordinates(address)
  if err != nil {
    t.Fatal(err)
  }
  if coords.Latitude != 55.751999 || coords.Longitude != 37.617734 {
    t.Fatalf("Coordinates for %s are wrong: %f, %f", address, coords.Latitude, coords.Longitude)
  }
}

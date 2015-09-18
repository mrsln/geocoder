/*
Package geocoder is a wrapper for the Yandex Geocoder API.

SUMMARY
  address := "Moscow, Kremlin"
  coords, err := geocoder.AddressToCoordinates(address)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Coordinates for '%s' are: %f (Latitude) %f (Longitude)\n",
    address, coords.Latitude, coords.Longitude)
*/
package geocoder

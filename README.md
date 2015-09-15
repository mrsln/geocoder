Yandex Geocoder API wrapper in GO
=================================
![travis](https://travis-ci.org/mrsln/geocoder.svg?branch=master)

The package will get coordinates for a given address.
```
address := "Moscow, Kremlin"
coords, err := geocoder.AddressToCoordinates(address)
if err != nil {
  panic(err)
}
fmt.Printf("Coordinates for '%s' are: %f (Latitude) %f (Longitude)\n",
  address, coords.Latitude, coords.Longitude)
```

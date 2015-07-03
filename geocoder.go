// Package geocoder encapsulates API for https://tech.yandex.ru/maps/doc/geocoder/desc/concepts/About-docpage/
package geocoder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Reply struct {
	Response Response
}

type Response struct {
	GeoObjectCollection GeoObjectCollection
}

type GeoObjectCollection struct {
	FeatureMember []FeatureMember
}

type FeatureMember struct {
	GeoObject GeoObject
}

type GeoObject struct {
	Point Point
}

type Point struct {
	Pos string
}

type Coordinates struct {
	Latitude  float32
	Longitude float32
}

func Request(address string) (Reply, error) {
	url := "http://geocode-maps.yandex.ru/1.x/?format=json&results=1&geocode=" + address
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return Reply{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Reply{}, err
	}
	var r Reply
	err = json.Unmarshal(body, &r)
	if err != nil {
		return Reply{}, err
	}
	return r, nil
}

func AddressToCoordinates(address string) (Coordinates, error) {
	r, err := Request(address)
	if err != nil {
		return Coordinates{}, err
	}
	if len(r.Response.GeoObjectCollection.FeatureMember) == 0 {
		return Coordinates{}, errors.New("geocoder: the place hasn't been found: " + address)
	}
	pos := r.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos
	points := strings.Split(pos, " ")
	if len(points) < 2 {
		return Coordinates{}, errors.New("geocoder: unexpected format of points (" + pos + ")")
	}
	lat, err := strconv.ParseFloat(points[1], 32)
	if err != nil {
		return Coordinates{}, err
	}
	lon, err := strconv.ParseFloat(points[0], 32)
	if err != nil {
		return Coordinates{}, err
	}
	c := Coordinates{float32(lat), float32(lon)}
	return c, nil
}

package main

import (
	"fmt"
	"math"
)

const (
	// EARTH_RADIUS https://www.google.com.vn/search?q=eath+radius&oq=eath+radius&aqs=chrome..69i57j0l5.2911j0j4&sourceid=chrome&ie=UTF-8
	EARTH_RADIUS = 6371
)

func degreeToRadian(d float64) float64 {
	return d * math.Pi / 180.0
}

// Coordinate defines struct of a coordinate
type Coordinate struct {
	Lat float64
	Lng float64
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%v,%v)", c.Lat, c.Lng)
}

// GreatCircleDistance uses formula from https://en.wikipedia.org/wiki/Great-circle_distance to calculate spherical
// distance from one point to another
func (p Coordinate) GreatCircleDistance(p2 Coordinate) float64 {
	dLat := degreeToRadian(p2.Lat - p.Lat)
	dLng := degreeToRadian(p2.Lng - p.Lng)

	a1 := math.Sin(dLat/2.0) * math.Sin(dLat/2.0)
	a2 := math.Cos(degreeToRadian(p.Lat)) * math.Cos(degreeToRadian(p2.Lat)) * math.Pow(math.Sin(dLng/2.0), 2)

	c := 2 * math.Asin(math.Sqrt(a1+a2))

	return EARTH_RADIUS * c
}

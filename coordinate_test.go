package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegreeToRadian(t *testing.T) {
	assert.Equal(t, 90*math.Pi/180.0, degreeToRadian(90))
}

func TestGreatCircleDistance(t *testing.T) {
	c1 := Coordinate{
		Lat: 51.925146,
		Lng: 4.478617,
	}

	c2 := Coordinate{
		Lat: 37.1768672,
		Lng: -3.608897,
	}

	// Expected result from http://www.onlineconversion.com/map_greatcircle_distance.htm
	assert.Equal(t, 1758.0806131200134, c2.GreatCircleDistance(c1))
}

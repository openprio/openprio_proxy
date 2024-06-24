package main

import (
	"math"
	"openprio_proxy/openprio_pt_position_data"
)

func BearingBetweenPositions(coord1 openprio_pt_position_data.Position, coord2 openprio_pt_position_data.Position, hdop float32) float32 {
	// make sure to set bearing on unknown when there is no good GPS measurement.
	if hdop >= 10.0 {
		return -1
	}

	lat1 := degToRad(coord1.Latitude)
	lon1 := degToRad(coord1.Longitude)
	lat2 := degToRad(coord2.Latitude)
	lon2 := degToRad(coord2.Longitude)

	dLon := float64(lon2 - lon1)

	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) -
		math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)

	bearing := math.Atan2(y, x)
	bearing = radToDeg(float32(bearing))
	if bearing < 0 {
		bearing = bearing + 360.0
	}
	return float32(bearing)
}

func degToRad(deg float32) float64 {
	return float64(deg) * (math.Pi / 180)
}

func radToDeg(rad float32) float64 {
	return float64(rad) * (180 / math.Pi)
}

func flatEarthDistance(lat1, lon1, lat2, lon2 float32) float32 {
	const R = 6371000 // Radius of the Earth in meters

	// Convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Average latitude for x-axis correction
	avgLat := (lat1Rad + lat2Rad) / 2

	// Calculate deltas
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Convert deltas to meters
	x := float64(deltaLon) * math.Cos(float64(avgLat)) * R
	y := float64(deltaLat * R)

	// Pythagorean theorem to calculate the distance
	distance := math.Sqrt(x*x + y*y)

	return float32(distance)
}

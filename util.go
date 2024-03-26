package main

import (
	"math"
	"openprio_proxy/openprio_pt_position_data"
)

func BearingBetweenPositions(coord1 openprio_pt_position_data.Position, coord2 openprio_pt_position_data.Position) float32 {
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

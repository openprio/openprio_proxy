package main

import (
	"testing"
)

func TestFleatEarthDistance1(t *testing.T) {
	// Example coordinates
	lat1, lon1 := float32(52.0168838500977), float32(4.35196542739868)
	lat2, lon2 := float32(52.0168838500977), float32(4.35195970535278)

	got := flatEarthDistance(lat1, lon1, lat2, lon2)
	if got > 1.5 {
		t.Errorf("flatEarthDistance = %f; want 1", got)
	}
}

func TestFleatEarthDistance2(t *testing.T) {
	// Example coordinates
	lat1, lon1 := float32(52.01687241), float32(4.351946831)
	lat2, lon2 := float32(52.01686478), float32(4.351935863)

	got := flatEarthDistance(lat1, lon1, lat2, lon2)
	if got > 1.5 {
		t.Errorf("flatEarthDistance = %f; want 1", got)
	}
}

func TestFleatEarthDistance3(t *testing.T) {
	// Example coordinates
	lat1, lon1 := float32(52.01687241), float32(4.351946831)
	lat2, lon2 := float32(52.01684952), float32(4.351914883)

	got := flatEarthDistance(lat1, lon1, lat2, lon2)
	if got < 1.5 {
		t.Errorf("flatEarthDistance = %f; want > 1.5", got)
	}
}

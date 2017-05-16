package intersect

import (
	"github.com/paulmach/go.geo"
	"testing"
)

func TestIntersect(t *testing.T) {
	// gps path
	gpslog := []GPSLog{{timestamp: 1494919104.0, longitude: 30, latitude: 45},
		{timestamp: 1494919105.0, longitude: 60, latitude: 45},
		{timestamp: 1494919106.0, longitude: 90, latitude: 45}}

	// line
	line := geo.NewLine(geo.NewPoint(40, 0), geo.NewPoint(40, 90))

	timestamp, point := Intersect(gpslog, line)
	if timestamp != 1494919104.0 {
		t.Errorf("Timestamp should be 0")
	}
	if point.Lng() != float64(40.0) {
		t.Errorf("Longitude should be 40.0 (lng=%f)", point.Lng())
	}
	if point.Lat() != float64(45.0) {
		t.Errorf("Latitude should be 45.0 (lat=%f)", point.Lat())
	}
}

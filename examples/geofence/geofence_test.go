package geofence

import (
	"github.com/kellydunn/golang-geo"
	"testing"
)

func TestGeoFence(t *testing.T) {
	var in_or_out Direction = NOTHING
	fence := geo.NewPolygon([]*geo.Point{geo.NewPoint(5, 5), geo.NewPoint(10, 5), geo.NewPoint(10, 10), geo.NewPoint(5, 10)})

	// in
	gpslog := [2]GPSLog{{timestamp: 1494919104.0, longitude: 0, latitude: 5},
		{timestamp: 1494919105.0, longitude: 7, latitude: 8}}
	if GeoFence(gpslog, fence) != IN {
		t.Errorf("in_or_out should be 1 (in_out=%d)", in_or_out)
	}

	// out
	gpslog = [2]GPSLog{{timestamp: 1494919105.0, longitude: 7, latitude: 8},
		{timestamp: 1494919104.0, longitude: 0, latitude: 5}}
	if GeoFence(gpslog, fence) != OUT {
		t.Errorf("in_or_out should be 2 (in_out=%d)", in_or_out)
	}

	// nothing
	gpslog = [2]GPSLog{{timestamp: 1494919105.0, longitude: 0, latitude: 8},
		{timestamp: 1494919104.0, longitude: 0, latitude: 5}}
	if GeoFence(gpslog, fence) != NOTHING {
		t.Errorf("in_or_out should be 0 (in_out=%d)", in_or_out)
	}
}

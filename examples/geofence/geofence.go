package geofence

import (
	"github.com/kellydunn/golang-geo"
)

type Direction int

const (
	NOTHING = iota
	IN
	OUT
)

func GeoFence(gpslog [2]GPSLog, fence *geo.Polygon) Direction {
	p_1 := fence.Contains(geo.NewPoint(gpslog[0].latitude, gpslog[0].longitude))
	p_2 := fence.Contains(geo.NewPoint(gpslog[1].latitude, gpslog[1].longitude))

	if p_1 == false && p_2 == true {
		return IN
	} else if p_1 == true && p_2 == false {
		return OUT
	}

	return NOTHING
}

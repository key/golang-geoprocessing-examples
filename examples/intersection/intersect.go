package intersect

import "github.com/paulmach/go.geo"

func Intersect(gpslog []GPSLog, line *geo.Line) (int, *geo.Point) {
	path := geo.NewPath()
	for _, x := range gpslog {
		path.Push(geo.NewPoint(x.longitude, x.latitude))
	}

	if path.Intersects(line) {
		points, _ := path.Intersection(line)
		for i, point := range points {
			return gpslog[i].timestamp, point
		}
	}

	return -1, nil
}

# GPS log intersection example

Detects that GPS log passed through specified line.

## Depended packages

- [go.geo](https://github.com/paulmach/go.geo)

## Usage

```go
gpslog := []GPSLog{{timestamp: 1494919104.0, longitude: 30, latitude: 45},
	{timestamp: 1494919105.0, longitude: 60, latitude: 45},
	{timestamp: 1494919106.0, longitude: 90, latitude: 45}}

line := geo.NewLine(geo.NewPoint(40, 0), geo.NewPoint(40, 90))

timestamp, point = Intersect(gpslog, line)
```

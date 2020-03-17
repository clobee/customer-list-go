package calculation

import (
	"github.com/umahmood/haversine"
)

const oLat = 53.339428
const oLong = -6.257664

func GetDistance(lat float64, lon float64) float64 {
	r := haversine.Coord{Lat: lat, Lon: lon}
	o := haversine.Coord{Lat: oLat, Lon: oLong}
	_, km := haversine.Distance(r, o)
	return km
}

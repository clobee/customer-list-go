package calculation

import (
	"github.com/umahmood/haversine"
)

const oLong = 53.339428
const oLat = -6.257664

func GetCoord(lat float64, lon float64) float64 {
	r := haversine.Coord{Lat: lat, Lon: lon}
	o := haversine.Coord{Lat: oLat, Lon: oLong}
	_, km := haversine.Distance(r, o)
	return km
}

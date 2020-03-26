package calculation

import (
	"github.com/clobee/customer-list-go/config"
	"github.com/umahmood/haversine"
)

func GetDistance(lat float64, lon float64) float64 {
	config := config.GetConf("/config.yaml")

	r := haversine.Coord{Lat: lat, Lon: lon}
	o := haversine.Coord{Lat: config.OLat, Lon: config.OLong}
	_, km := haversine.Distance(r, o)
	return km
}

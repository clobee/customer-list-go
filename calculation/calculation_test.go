package calculation

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	fixtures := []struct {
		latitude  float64
		longitude float64
		result    float64
	}{
		{53.339428, -6.257664, 0},
		{53.1489345, -6.257664, 21.181910758765177},
		{53.008769, -6.1056711, 38.13756809820155},
	}

	for _, data := range fixtures {
		co := GetDistance(data.latitude, data.longitude)

		if data.result != co {
			t.Error(
				"For latitude:", data.latitude,
				"& longitude:", data.longitude,
				"expected:", data.result,
				"got:", co,
			)
		}
	}
}

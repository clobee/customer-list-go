package calculation

import "testing"

func TestGetDistance(t *testing.T) {
	fixtures := []struct {
		latitude  float64
		longitude float64
		result    float64
	}{
		{53.339428, -6.257664, 0},
		{53.1489345, -6.257664, 21.181911},
		{53.008769, -6.1056711, 38.137568},
	}

	for _, data := range fixtures {
		co := GetDistance(data.latitude, data.longitude)

		if data.result != co {
			t.Errorf("Incorrect distance for lat %f and long %f. Expected: %f, Found: %f", data.latitude, data.longitude, data.result, co)
		}
	}
}

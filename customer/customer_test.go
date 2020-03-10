package customer

import (
	"reflect"
	"testing"
)

func TestGetInfoCreatesCustomers(t *testing.T) {

	fixtures := []string{
		`{"latitude": "53.761389", "user_id": 30, "name": "Nick Enright", "longitude": "-7.2875"}`,
		`{"latitude": "54.080556", "user_id": 23, "name": "Eoin Gallagher", "longitude": "-6.361944"}`,
		`{"latitude": "52.833502", "user_id": 25, "name": "David Behan", "longitude" : "-6.361944"}`,
	}

	for i := range fixtures {
		c := GetInfo(fixtures[i])

		s := reflect.TypeOf(c.Latitude).String()
		if "string" != s {
			t.Errorf("Latitude should be a string. Founded: %s", s)
		}

		v := reflect.TypeOf(c.UserID).String()
		if "int" != v {
			t.Errorf("UserID should be a int. Founded: %s", v)
		}

		w := reflect.TypeOf(c.Name).String()
		if "string" != w {
			t.Errorf("Name should be a string. Founded: %s", w)
		}

		h := reflect.TypeOf(c.Longitude).String()
		if "string" != h {
			t.Errorf("Longitude should be a string. Founded: %s", h)
		}
	}
}

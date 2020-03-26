package config

import (
	"reflect"
	"testing"
)

func TestGetConf(t *testing.T) {
	config := GetConf("../config.yaml")

	var s = ""

	s = reflect.TypeOf(config.OLat).String()
	if "float64" != s {
		t.Error(
			"Issue with config.OLat.",
			"expected: float64",
			"got:", s,
		)
	}

	s = reflect.TypeOf(config.OLong).String()
	if "float64" != s {
		t.Error(
			"Issue with config.oLong.",
			"expected: float64",
			"got:", s,
		)
	}

	s = reflect.TypeOf(config.LimitDistance).String()
	if "float64" != s {
		t.Error(
			"Issue with limitDistance.",
			"expected: float64",
			"got:", s,
		)
	}

	s = reflect.TypeOf(config.CustomerDataFile).String()
	if "string" != s {
		t.Error(
			"Issue with customerDataFile.",
			"expected: string",
			"got:", s,
		)
	}
}

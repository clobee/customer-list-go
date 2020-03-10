package data

import (
	"encoding/json"
	"testing"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func TestReadFileReturnsJsonStrings(t *testing.T) {

	content := ReadFile("../fixtures/customers.txt")

	for z := range content {

		if true != isJSON(z) {
			t.Errorf("This content is not a JSON String: %v", z)
		}
	}
}

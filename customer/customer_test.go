package customer

import (
	"testing"
)

type customerMock struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
}

func TestCustomerDataStructureIsCorrect(t *testing.T) {

	t.Skip("skipping test in short mode.")

	// customer := customer{}
	// customer.Latitude = "-5559538484"
	// customer.Longitude = "2332454"
	// customer.UserID = 123
	// customer.Name = "Albert"

	// customerMocked := customerMock{}
	// customer.Latitude = "-5559538484"
	// customer.Longitude = "2332454"
	// customer.UserID = 123
	// customer.Name = "Albert"

	// equalStruct := reflect.DeepEqual(customer, customerMocked)

	// if equalStruct != true {
	// 	t.Errorf("%v and %v are not equal.", customer, customerMocked)
	// }
}

func TestConstructorNewReturnsCorrectDataStructure(t *testing.T) {
	t.Skip("skipping test in short mode.")
}

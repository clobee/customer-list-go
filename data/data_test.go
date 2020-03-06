package data

// import (
// 	"bufio"
// 	"os"
// 	// "reflect"
// 	"github.com/clobee/customer-list-go/customer"
// 	"testing"
// )

// func TestDataProcessJson(t *testing.T) {

// 	expectedType := customer{}

// 	file, err := os.Open("../fixtures/customers.txt")

// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		data := processLine(scanner.Text())

// 		// equalStruct := reflect.DeepEqual(customer, customerMocked)

// 		// if equalStruct != true {
// 		t.Errorf("%T type vs %T.", data, expectedType)
// 		// }
// 	}

// 	if err != nil {
// 		t.Errorf("%s", scanner.Err())
// 	}

// }

// func TestCustomerDataStructureIsCorrectX(t *testing.T) {

// 	t.Skip("skipping test in short mode.")

// }

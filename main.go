package main

import (
	"fmt"
	"github.com/clobee/customer-list-go/data"
)

func main() {
	results := data.ReadFile("fixtures/customers.txt")

	for line := range results {
		fmt.Printf("customer : '%v'\n", line)
	}
}

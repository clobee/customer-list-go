package main

import (
	"fmt"
	"github.com/clobee/customer-list-go/customer"
	"github.com/clobee/customer-list-go/data"
)

func main() {
	content := data.ReadFile("fixtures/customers.txt")

	for line := range content {
		user := customer.GetInfo(line)
		fmt.Printf("customer : '%v' (%d)\n", user.Name, user.UserID)
	}
}

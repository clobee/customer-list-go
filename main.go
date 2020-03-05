package main

import (
	"github.com/clobee/customer-list-go/data"
)

func main() {
	data.ReadFile("fixtures/customers.txt")
}

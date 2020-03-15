package main

import (
	"fmt"
	"github.com/clobee/customer-list-go/calculation"
	"github.com/clobee/customer-list-go/customer"
	"github.com/clobee/customer-list-go/data"
	"log"
	"strconv"
)

const limitDistance = 100

func main() {
	content := data.ReadFile("fixtures/customers.txt")

	selCustomers := make(map[int]customer.Customer)

	for line := range content {
		u := customer.GetInfo(line)

		l, err := strconv.ParseFloat(u.Longitude, 64)
		la, err := strconv.ParseFloat(u.Latitude, 64)

		if err != nil {
			log.Fatal(err)
		}

		c := calculation.GetCoord(l, la)
		if c < limitDistance {
			selCustomers[u.UserID] = u
		}
	}

	for key, value := range selCustomers {
		fmt.Printf("%d : %s \n", key, value.Name)
	}
}

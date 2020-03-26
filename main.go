package main

import (
	"fmt"
	"github.com/clobee/customer-list-go/calculation"
	"github.com/clobee/customer-list-go/config"
	"github.com/clobee/customer-list-go/customer"
	"github.com/clobee/customer-list-go/data"
	"log"
	"sort"
	"strconv"
)

const configFile string = "config.yaml"

type ByUserID []customer.Customer

func (a ByUserID) Len() int           { return len(a) }
func (a ByUserID) Less(i, j int) bool { return a[i].UserID < a[j].UserID }
func (a ByUserID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	config := config.GetConf(configFile)
	content := data.ReadFile(config.CustomerDataFile)

	selCustomers := make([]customer.Customer, 0, 1)

	for line := range content {
		u := customer.GetInfo(line)

		lo, err := strconv.ParseFloat(u.Longitude, 64)
		la, err := strconv.ParseFloat(u.Latitude, 64)

		if err != nil {
			log.Fatal(err)
		}

		c := calculation.GetDistance(la, lo, configFile)
		if c < config.LimitDistance {
			selCustomers = append(selCustomers, u)
		}
	}

	sort.Sort(ByUserID(selCustomers))

	for key, value := range selCustomers {
		fmt.Printf("%d : %s %d \n", key, value.Name, value.UserID)
	}
}

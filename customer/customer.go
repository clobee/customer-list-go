package customer

import (
	"encoding/json"
	"log"
)

type Customer struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
}

//GetInfo takes a Json string and returns a customer
func GetInfo(msg string) Customer {
	data := Customer{}

	if err := json.Unmarshal([]byte(msg), &data); err != nil {
		log.Fatal(err)
	}

	return data
}

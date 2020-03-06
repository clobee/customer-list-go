package data

import (
	"bufio"
	"encoding/json"
	"github.com/clobee/customer-list-go/customer"
	"log"
	"os"
)

func ReadFile(document string) chan customer.Customer {
	file, err := os.Open(document)

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan customer.Customer)

	go func() {
		defer file.Close()
		defer close(results)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			results <- processLine(scanner.Text())
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	return results
}

func processLine(msg string) customer.Customer {
	data := customer.Customer{}

	if err := json.Unmarshal([]byte(msg), &data); err != nil {
		log.Fatal(err)
	}

	return data
}

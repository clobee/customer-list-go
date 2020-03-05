package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/clobee/customer-list-go/customer"
	"log"
	"os"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(document string) {
	file, err := os.Open(document)

	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)

		go func(msg string) {
			defer wg.Done()
			customerData := processLine(msg)
			fmt.Printf("customer : '%v'\n", customerData)
		}(scanner.Text())
	}

	wg.Wait()
	check(scanner.Err())
}

func processLine(msg string) customer.Customer {
	data := customer.Customer{}

	if err := json.Unmarshal([]byte(msg), &data); err != nil {
		log.Fatal(err)
	}

	return data
}

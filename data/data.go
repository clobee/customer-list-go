package data

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(document string) chan string {
	file, err := os.Open(document)

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan string)

	go func() {
		defer file.Close()
		defer close(results)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// msg := scanner.Text()
			results <- scanner.Text()
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	return results
}

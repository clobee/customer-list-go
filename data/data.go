package data

import (
	"bufio"
	"log"
	"os"
)

//ReadFile receive a json file and returns a channel with each line
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
			results <- scanner.Text()
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	return results
}

# The Test

### 1. Technical problem

We have some customer records in a text file (customers.txt) -- one customer per line, JSON lines formatted. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).

    You must use the first formula from this article https://en.wikipedia.org/wiki/Great-circle_distance to calculate distance. Don't forget, you'll need to convert degrees to radians.

    The GPS coordinates for our Dublin office are 53.339428, -6.257664.

    You can find the Customer list here.

### Application

run tests : `go test -v ./tests`
run the app: `go run *.go`
create a build: `go build` //Which doesn't includes tests



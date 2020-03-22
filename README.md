# The Test

[![BCH compliance](https://bettercodehub.com/edge/badge/clobee/customer-list-go?branch=master)](https://bettercodehub.com/)
[![CircleCI](https://circleci.com/gh/clobee/customer-list-go.svg?style=svg)](https://circleci.com/gh/clobee/customer-list-go)

### 1. Technical problem

We have some customer records in a text file (customers.txt) -- one customer per line, JSON lines formatted. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).

    You must use the first formula from this article https://en.wikipedia.org/wiki/Great-circle_distance to calculate distance. 

    The GPS coordinates for our office are 53.339428, -6.257664.

### Application

test the program: `make test`
run the program: `make run`




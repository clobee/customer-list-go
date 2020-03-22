default: test

test:
	cd ./src ; \
	go test -v ./... 2>&1 ; \
	
run: 
	go build; go install; $HOME/go/bin/customer-list-go

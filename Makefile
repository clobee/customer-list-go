default: test

test:
	cd ./src ; \
	go test -v ./... 2>&1 ; \

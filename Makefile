MAKEFLAGS += --silent
APP=ipaddr

default: run

.PHONY: install
## install: install and compile package dependencies
install: tidy
	go install

.PHONY: tidy
## tidy: add missing and remove unused modules
tidy:
	go mod tidy

.PHONY: run
## run: build and run the cli tool
run: build
	./${APP}

.PHONY: build
## build: build the cli tool
build: clean install
	go build -ldflags="-s -w -X 'github.com/adityasanka/ipaddr.Token=${TOKEN}'" -o ${APP} ./cmd

.PHONY: clean
## clean: remove binaries
clean:
	rm -f ${APP}
	go clean

.PHONY: help
## help: prints this help message
help:
	echo "Usage: \n"
	sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
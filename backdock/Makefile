SHELL := /bin/bash

all: build

dev:
	source ./env.sh && go run ./main.go

build:
	go build -o bin/main ./main.go

clean:
	rm ./bin/main

PWD=$(shell pwd)
BINARY_NAME=cvapi
ENV_FILE=.exportme

build:
	go build -o $(BINARY_NAME) .
run:
	. $(PWD)/$(ENV_FILE) && ./$(BINARY_NAME)
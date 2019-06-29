SRC = $(wildcard *.go)

.PHONY: all
all: build


.PHONY: build
build: grvm

grvm: ${SRC}
	go build -o grvm

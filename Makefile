SRC=$(wildcard *.go)
DISTFILE=grvm-amd64.tar.gz

.PHONY: all
all: build

dist:
	mkdir -p dist

.PHONY: build
build: dist/grvm

dist/grvm: ${SRC}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o dist/grvm *.go

.PHONY: package
package: dist/grvm
	cp README.md dist/
	cp LICENSE.txt dist/
	cd dist && tar -czvf ${DISTFILE} *

.PHONY: clean
clean:
	rm -rf dist
	rm ${DISTFILE}

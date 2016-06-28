.PHONY: tests

default: tests

deps:
	go get -u github.com/Sirupsen/logrus
	go get -u github.com/stretchr/testify/assert

fmt:
	go fmt github.com/gtaylor/dmarc/...

tests:
	go test github.com/gtaylor/dmarc/tests


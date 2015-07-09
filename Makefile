.PHONY: build clean

BASE=${PWD}
TARGET := webbase.goc

GOPATH := ${PWD}/:${GOPATH}

export GOPATH

default: build

build:
	@echo ${GOPATH}
	go build -o ${BASE}/${TARGET} ./src/main.go

clean:
	@rm -rf ${BASE}/${TARGET}
	@echo "clean ok"
# Makefile for the Docker image stevesloka/webapp-healthz
# MAINTAINER: Steve Sloka <steve@stevesloka.com>
# If you update this image please bump the tag value before pushing.

.PHONY: all build container push clean test

TAG = 1.0.0
PREFIX = 192.168.2.50:5000/stevesloka

all: container

build: main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o webapp-healthz --ldflags '-w' ./main.go

container: build
	docker build -t $(PREFIX)/webapp-healthz:$(TAG) .

push:
	docker push $(PREFIX)/webapp-healthz:$(TAG)

clean:
	rm -f restapi

test: clean
	godep go test -v --vmodule=*=4

ARCH ?= linux

all: build

build: build/opspki

build/opspki: main.go
	CGO_ENABLED=0 GOOS=$(ARCH) go build -a -tags netgo -ldflags '-w' -o build/opspki -v .

clean:
	rm -rvf build
	rm -vf docker/opspki

mrproper: clean
	rm -rvf vendor

docker: build
	cp -v build/opspki docker/
	cd docker && \
	docker build -t opspki .

.PHONY: all build clean docker
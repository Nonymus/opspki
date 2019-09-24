all: build

build: build/toldyouso

build/toldyouso: main.go
	CGO_ENABLED=0 go build -o build/toldyouso -v .

clean:
	rm -rvf build

.PHONY: all build clean
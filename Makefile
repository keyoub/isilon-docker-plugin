PACKAGES := \
	github.west.isilon.com/bkeyoumarsi/docker-plugin \
	github.west.isilon.com/bkeyoumarsi/docker-plugin/driver
DEPENDENCIES := github.com/calavera/dkvolume

all: build silent-test

build:
	go build -o plugin

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)

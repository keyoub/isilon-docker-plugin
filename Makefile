PACKAGES := \
	github.west.isilon.com/bkeyoumarsi/docker-plugin \
	github.west.isilon.com/bkeyoumarsi/docker-plugin/driver
DEPENDENCIES := github.com/calavera/dkvolume

install: deps
	go install -o isi-plugin

build:
	go build -o isi-plugin

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)

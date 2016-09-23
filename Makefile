PREFIX?=/build

GOOS = linux
GOARCH = amd64

GOFILES = $(shell find . -type f -name '*.go')
uwsgibeat: $(GOFILES)
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm uwsgibeat || true

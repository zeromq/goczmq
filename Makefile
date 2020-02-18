build: format clean test
	go build ./...

test: get
	go test -v .

bench: get
	go test -v -bench . ./...

get:
	go get -t -v ./...

format:
	find . -name \*.go -type f -exec gofmt -w {} \;

perf:
	$(MAKE) -C ./cmd/perf build

clean:
	$(MAKE) -C ./cmd/perf clean

.PHONY: clean build

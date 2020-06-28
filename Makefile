.PHONY: clean lint test build/check-exit

clean:
	rm -rf ./build

lint:
	golangci-lint run --config ./.golangci.yml

test:
	go test -v .

build/check-exit:
	go build -ldflags="-s -w" -o ./build/check-exit ./cmd/check-exit

build: build/check-exit
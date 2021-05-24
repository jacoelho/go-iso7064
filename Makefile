default: test

test:
	go test -race -v ./...

ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }


escape:
	go build -gcflags='-m' ./...

.PHONY: ci-tidy escape test

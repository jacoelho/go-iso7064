default: test

test:
	go test -race -v ./...

ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }


escape:
	go build -gcflags='-m' ./...

bounds:
	go build -gcflags="-d=ssa/check_bce/debug=3" ./...

.PHONY: ci-tidy escape test bounds

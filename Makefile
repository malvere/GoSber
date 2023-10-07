.PHONY: build
build:
	go build -v -o sber-scrape ./cmd/sber-scrape

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: prod
prod:
	@if [ "$(filter windows,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=windows GOARCH=amd64 go build -o sber-scrape-win-x86.exe -v ./cmd/sber-scrape; \
	elif [ "$(filter macos,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o sber-scrape-darwin-amd64 -v ./cmd/sber-scrape; \
	elif [ "$(filter linux-386,$(MAKECMDGOALS))" != "" ]; \
		GOOS=linux GOARCH=386 go build -o sber-scrape-linux-386 -v ./cmd/sber-scrape; \
	else \
		GOOS=linux GOARCH=amd64 go build -o sber-scrape-linux-amd64 -v ./cmd/sber-scrape; \
	fi

PHONY: clean
clean:
	rm -f sber-scrape
	
.DEFAULT_GOAL := build

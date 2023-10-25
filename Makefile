.PHONY: build
build:
	go build -v -o ./bin/sber-scrape ./cmd/sber-scrape

.PHONY: prod
prod:
	@if [ "$(filter windows,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=windows GOARCH=amd64 go build -o ./bin/sber-scrape-win-x86.exe -v ./cmd/sber-scrape; \
	elif [ "$(filter macos,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o ./bin/sber-scrape-darwin-amd64 -v ./cmd/sber-scrape; \
	elif [ "$(filter linux-386,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=linux GOARCH=386 go build -o ./bin/sber-scrape-linux-386 -v ./cmd/sber-scrape; \
	else \
		GOOS=linux GOARCH=amd64 go build -o ./bin/sber-scrape-linux-amd64 -v ./cmd/sber-scrape; \
	fi

.PHONY: run

run:
	go run ./cmd/sber-scrape/main.go

.PHONY: clean
clean:
	rm -f /bin/

.DEFAULT_GOAL := build
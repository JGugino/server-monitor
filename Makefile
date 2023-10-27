build:
	@go build -o bin/smonitor

run: build
	@./bin/smonitor
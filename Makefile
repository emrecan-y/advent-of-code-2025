SHELL := /usr/bin/sh
DAY=00

.PHONY: run
run:
	@if [ -f ./cmd/day${DAY}/main.go ]; then \
		time go run ./cmd/day${DAY}/main.go; \
	else \
		echo "ERROR: The implementation for day ${DAY} could not be found."; \
	fi
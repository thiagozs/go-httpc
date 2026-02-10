default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: lint
lint: # Run linter
	# Run linter
	# ----------
	@golangci-lint run
	# Lint complete +++

.PHONY: test
test: # Run tests
	# Run tests
	# ---------
	@go test -coverpkg=./... -coverprofile=coverage.out ./... && \
	grep -v -E "_mock.go|.sql.go|db.go" coverage.out > filtered_coverage.out && \
	go tool cover -func filtered_coverage.out && \
	rm coverage.out filtered_coverage.out
	# Tests complete +++

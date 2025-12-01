.PHONY: benchmark test

benchmark:
	@go test -run=^$$ -bench=. -benchmem ./...

test:
	@go test -cover ./...
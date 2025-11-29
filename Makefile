.PHONY: test

test:
	go test ./... -cover -bench=. -benchmem
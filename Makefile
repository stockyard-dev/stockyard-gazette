build:
	CGO_ENABLED=0 go build -o gazette ./cmd/gazette/

run: build
	./gazette

test:
	go test ./...

clean:
	rm -f gazette

.PHONY: build run test clean

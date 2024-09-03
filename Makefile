build:
  @go build -o bin/ecom_api cmd/main.go

test:
  @go test -v ./...

run: build
  @./bin/ecom_api
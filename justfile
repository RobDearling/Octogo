set windows-shell := ["pwsh", "-NoLogo", "-Command"]

default:
  just --list

fmt:
  gofmt -s -w -e .

tidy:
  go mod tidy

build:
  go build -o ./bin -v ./...

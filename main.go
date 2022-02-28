package main

import (
	"github.com/brittonhayes/pillager/cmd/pillager"
)

//go:generate golangci-lint run ./...
//go:generate gomarkdoc ./pkg/hunter/...
//go:generate gomarkdoc ./cmd/pillager/...

func main() {
	pillager.Execute()
}

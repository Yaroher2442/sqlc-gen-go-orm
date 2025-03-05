package main

import (
	"github.com/sqlc-dev/plugin-sdk-go/codegen"

	golang "github.com/Yaroher2442/sqlc-gen-go-orm/internal"
)

func main() {
	codegen.Run(golang.Generate)
}

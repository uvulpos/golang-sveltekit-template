package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/uvulpos/go-svelte/src/cmd"
	_ "github.com/uvulpos/go-svelte/swagger-docs"
)

func main() {
	cmd.Execute()
}

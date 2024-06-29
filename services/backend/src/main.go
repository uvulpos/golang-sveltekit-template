package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/uvulpos/go-svelte/src/cmd"
	"github.com/uvulpos/go-svelte/src/configuration"
	_ "github.com/uvulpos/go-svelte/swagger-docs"
)

func main() {
	configuration.LoadConfiguation()
	cmd.Execute()
}

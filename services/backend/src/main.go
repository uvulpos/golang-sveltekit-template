package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/uvulpos/golang-sveltekit-template/src/cmd"
	_ "github.com/uvulpos/golang-sveltekit-template/swagger-docs"
)

func main() {
	cmd.Execute()
}

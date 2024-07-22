package assets

import "embed"

//go:embed frontend/*
var SvelteFS embed.FS

//go:embed terminal-assets/*
var TerminalFS embed.FS

//go:embed swagger-css/gopher-coffee.gif
var GopherCoffee string

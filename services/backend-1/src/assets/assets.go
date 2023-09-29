package assets

import "embed"

//go:embed frontend/*
var SvelteFS embed.FS

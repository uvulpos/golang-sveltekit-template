//go:build tools
// +build tools

package tools

import (
	_ "github.com/cortesi/modd/cmd/modd"
	_ "github.com/cortesi/termlog"
	_ "golang.org/x/tools/cmd/goimports"
)

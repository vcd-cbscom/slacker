// Package tools imports the Go binaries we use, it's used to version dependencies.
package tools

import (
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "honnef.co/go/tools/cmd/staticcheck"
)

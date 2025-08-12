package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/calimapp/x/gowork/workspace"
)

var getUsage func() = func() {
	fmt.Print(`
Usage:
    gowork get [module]

Example:
	gowork get github.com/foo/bar@latest
`)
}

func RunGet(args []string) {
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	fs.Usage = getUsage
	fs.Parse(args)

	if fs.NArg() != 1 {
		getUsage()
		os.Exit(1)
	}

	dependency := fs.Arg(0)

	ws := workspace.Read()

	for _, mod := range ws.Modules {
		workspace.GetDependency(mod, dependency)
	}
	ws.Dependencies = append(ws.Dependencies, dependency)
	ws.Save()
}

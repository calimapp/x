package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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
		absPath, _ := filepath.Abs(mod)
		fmt.Printf("🔄 Running go get %s in %s\n", dependency, absPath)

		// Run go get
		getCmd := exec.Command("go", "get", dependency)
		getCmd.Dir = absPath
		getCmd.Stdout = io.Discard
		getCmd.Stderr = io.Discard
		if err := getCmd.Run(); err != nil {
			log.Printf("❌ go get failed in %s: %v", mod, err)
			continue
		}

		// Run go mod tidy
		tidyCmd := exec.Command("go", "mod", "tidy")
		tidyCmd.Dir = absPath
		tidyCmd.Stdout = io.Discard
		tidyCmd.Stderr = io.Discard
		if err := tidyCmd.Run(); err != nil {
			log.Printf("❌ go mod tidy failed in %s: %v", mod, err)
		} else {
			fmt.Printf("✅ Synced %s in %s\n", dependency, mod)
		}
	}
	ws.Dependencies = append(ws.Dependencies, dependency)
	ws.Save()
}

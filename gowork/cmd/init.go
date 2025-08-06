package cmd

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/calimapp/x/gowork/workspace"
)

var initUsage func() = func() {
	fmt.Print(`
Usage:
    go work init [modules...]

Example:
	go work init ./microservice-a ./microservice-b ./common-lib
`)
}

func RunInit(args []string) {
	fs := flag.NewFlagSet("init", flag.ExitOnError)
	fs.Usage = initUsage
	fs.Parse(args)
	if workspace.WorkspaceExists() {
		log.Fatalf("Workspace file already exists, delete manually before recreate")
	}

	goVersion, err := detectGoVersion()
	if err != nil {
		log.Fatalf("Failed to detect Go version: %v", err)
	}

	ws := workspace.Workspace{
		Go:      goVersion,
		Modules: fs.Args(),
	}

	ws.Save()
	if len(ws.Modules) == 0 {
		fmt.Println("Workspace initialized")
	} else {
		fmt.Printf("Workspace initialized with modules: %q\n", ws.Modules)
	}
}

func detectGoVersion() (string, error) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		return "", err
	}

	parts := strings.Fields(string(out))
	if len(parts) < 3 {
		return "", fmt.Errorf("unexpected output from 'go version'")
	}

	version := strings.TrimPrefix(parts[2], "go")
	return version, nil
}

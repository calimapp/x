package main

import (
	"fmt"
	"os"

	"github.com/calimapp/x/gowork/cmd"
)

const usage string = `
Usage:
    go work <command> [arguments]

The commands are:
	init        initialize workspace file

Use "gowork <command> -help" for more information about a command.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "init":
		cmd.RunInit(os.Args[2:])
	case "get":
		cmd.RunGet(os.Args[2:])
	case "list":
		cmd.RunList()
	// case "use":
	// 	handleUse(os.Args[2:])
	// case "list":
	// 	handleList()
	// case "build":
	// 	handleBuild()
	default:
		fmt.Println("Unknown command:", command)
		fmt.Print(usage)
		os.Exit(1)
	}
}

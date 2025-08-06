package cmd

import (
	"fmt"

	"github.com/calimapp/x/gowork/workspace"
)

func RunList() {
	ws := workspace.Read()
	fmt.Println("Modules in workspace:")
	for _, m := range ws.Modules {
		fmt.Println("-", m)
	}
}

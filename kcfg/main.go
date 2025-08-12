package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/calimapp/x/kcfg/cmd"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	args := os.Args[1]

	switch args {
	case "merge":
		if err := cmd.MergeConfigs(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "list":
		if err := listContexts(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Missing context name")
			os.Exit(1)
		}
		if err := cmd.UseContext(os.Args[2]); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  kubemgr merge              Merge all kubeconfigs into ~/.kube/config")
	fmt.Println("  kubemgr list               List all contexts in merged config")
	fmt.Println("  kcm use <context>      Set current context in merged config")
	fmt.Println("  kubemgr use <context>      Set current context in merged config")
}

// listContexts shows all contexts in ~/.kube/config
func listContexts() error {
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".kube", "config")

	cfg, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return err
	}

	fmt.Println("Available contexts:")
	for name := range cfg.Contexts {
		marker := ""
		if name == cfg.CurrentContext {
			marker = " (current)"
		}
		fmt.Printf("  %s%s\n", name, marker)
	}
	return nil
}

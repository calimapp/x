package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
)

// useContext sets the current context in ~/.kube/config
func UseContext(context string) error {
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".kube", "config")

	cfg, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return err
	}

	if _, exists := cfg.Contexts[context]; !exists {
		return fmt.Errorf("context %q not found", context)
	}

	cfg.CurrentContext = context
	if err := clientcmd.WriteToFile(*cfg, configPath); err != nil {
		return err
	}

	fmt.Println("Switched to context:", context)
	return nil
}

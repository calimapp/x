package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// mergeConfigs aggregates ~/.kube/kcfg/*.yaml into ~/.kube/config
func MergeConfigs() error {
	home := os.Getenv("HOME")
	kcfgDir := filepath.Join(home, ".kube", "kcfg")
	files, _ := filepath.Glob(filepath.Join(kcfgDir, "*.yaml"))

	merged := api.NewConfig()

	for _, file := range files {
		cfg, err := clientcmd.LoadFromFile(file)
		if err != nil {
			fmt.Printf("Error loading %s: %v\n", file, err)
			continue
		}

		for name, cluster := range cfg.Clusters {
			merged.Clusters[name] = cluster
		}
		for name, authInfo := range cfg.AuthInfos {
			merged.AuthInfos[name] = authInfo
		}
		for name, context := range cfg.Contexts {
			merged.Contexts[name] = context
		}
		if cfg.CurrentContext != "" && merged.CurrentContext == "" {
			merged.CurrentContext = cfg.CurrentContext
		}
	}

	outPath := filepath.Join(home, ".kube", "config")
	if err := clientcmd.WriteToFile(*merged, outPath); err != nil {
		return err
	}

	fmt.Println("Merged kubeconfig saved to", outPath)
	return nil
}

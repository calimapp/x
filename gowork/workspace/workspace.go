package workspace

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

const WorkspaceFile = ".gowork.yaml"

type Workspace struct {
	Go           string   `yaml:"go"`
	Modules      []string `yaml:"modules"`
	Dependencies []string `yaml:"dependencies"`
}

func Read() *Workspace {
	data, err := os.ReadFile(WorkspaceFile)
	if err != nil {
		log.Fatal("gowork.yaml not found. Did you run 'gowork init'?")
	}
	var ws Workspace
	if err := yaml.Unmarshal(data, &ws); err != nil {
		log.Fatal("Invalid workspace file:", err)
	}
	return &ws
}

func WorkspaceExists() bool {
	info, err := os.Stat(WorkspaceFile)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (w *Workspace) Save() {
	data, err := yaml.MarshalWithOptions(w, yaml.IndentSequence(true))
	if err != nil {
		log.Fatal("Failed to marshal workspace YAML:", err)
	}

	err = os.WriteFile(WorkspaceFile, data, 0644)
	if err != nil {
		log.Fatal("Failed to write workspace.yaml:", err)
	}
}

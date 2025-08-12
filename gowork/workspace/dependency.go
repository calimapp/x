package workspace

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type modBackup struct {
	Dir     string
	ModFile string
	SumFile string
}

func GetDependency(module string, dependency string) {
	absPath, _ := filepath.Abs(module)
	// 🔒 Backup go.mod and go.sum
	modFile := filepath.Join(absPath, "go.mod")
	sumFile := filepath.Join(absPath, "go.sum")
	backup := modBackup{
		Dir:     absPath,
		ModFile: modFile + ".bak",
		SumFile: sumFile + ".bak",
	}
	copyFile(modFile, backup.ModFile)
	copyFile(sumFile, backup.SumFile)

	// Run go get
	getCmd := exec.Command("go", "get", dependency)
	getCmd.Dir = absPath
	getCmd.Stdout = io.Discard
	getCmd.Stderr = io.Discard
	if err := getCmd.Run(); err != nil {
		log.Printf("❌ go get failed in %s: %v", module, err)
		log.Println("🔁 Rolling back changes...")
		rollbackChanges(backup)
		return
	}

	// Run go mod tidy
	// tidyCmd := exec.Command("go", "mod", "tidy")
	// tidyCmd.Dir = absPath
	// tidyCmd.Stdout = io.Discard
	// tidyCmd.Stderr = io.Discard
	// if err := tidyCmd.Run(); err != nil {
	// 	log.Printf("❌ go get failed in %s: %v", module, err)
	// 	log.Println("🔁 Rolling back changes...")
	// 	rollbackChanges(backup)
	// 	return
	// }
	// fmt.Printf("✅ Synced %s in %s\n", dependency, module)
	os.Remove(backup.ModFile)
	os.Remove(backup.SumFile)
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func rollbackChanges(backup modBackup) {
	log.Printf("↩️ Restoring %s", backup.Dir)
	copyFile(backup.ModFile, filepath.Join(backup.Dir, "go.mod"))
	copyFile(backup.SumFile, filepath.Join(backup.Dir, "go.sum"))
	os.Remove(backup.ModFile)
	os.Remove(backup.SumFile)
}

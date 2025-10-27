package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"
)

func ProjectPath(projectName string) (string, error) {
	dir, _ := os.Getwd()
	projectPath := filepath.Join(dir, projectName)
	err := os.MkdirAll(projectPath, 0755)
	if err != nil {
		return "", fmt.Errorf("could not create project folder: %v", err)
	}
	return projectPath, nil

}

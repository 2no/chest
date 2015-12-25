package utils

import (
	"os"
	"path/filepath"
)

func GetRootPath() string {
	path := os.Getenv("CHEST_PATH")
	if path == "" {
		path, _ = os.Getwd()
	}
	return path
}

func GetChestPath() string {
	return filepath.Join(GetRootPath(), "/.chest/")
}

func GetGitIgnorePath() string {
	return filepath.Join(GetRootPath(), "/.gitignore")
}

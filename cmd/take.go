package cmd

import (
	"fmt"
	"github.com/wakuworks/chest/utils"
	"os"
	"path/filepath"
)

func Take(files []string) {
	chestPath := utils.GetChestPath()
	rootPath := filepath.Dir(chestPath)

	if !utils.ExistsFile(chestPath) {
		fmt.Println("Chest not found")
		return
	}

	for _, file := range files {
		chestInFilePath := filepath.Join(chestPath, file)
		if !utils.ExistsFile(chestInFilePath) {
			fmt.Println("File not found: " + chestInFilePath)
			return
		}

		rootInFilePath := filepath.Join(rootPath, file)
		if utils.ExistsFile(rootInFilePath) {
			fmt.Println("File already exist in root: " + rootInFilePath)
			fmt.Println("diff:")
			utils.PrintFileDiff(chestInFilePath, rootInFilePath)
			return
		}

		os.Remove(rootInFilePath)
		os.Rename(chestInFilePath, rootInFilePath)
	}
}

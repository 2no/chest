package cmd

import (
	"../utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Open() {
	chestPath := utils.GetChestPath()

	if !utils.ExistsFile(chestPath) {
		fmt.Println("Chest not found")
		return
	}

	files, err := ioutil.ReadDir(chestPath)
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Println("Chest is blank")
		return
	}

	rootPath := filepath.Dir(chestPath)
	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(rootPath, fileName)
		if utils.ExistsFile(filePath) {
			fmt.Println("File already exist: " + fileName)
			return
		}

		if err := os.Symlink(filepath.Join(chestPath, fileName), filePath); err != nil {
			panic(err)
		}
	}
}

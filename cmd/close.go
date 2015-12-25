package cmd

import (
	"fmt"
	"github.com/wakuworks/chest/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Close() {
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
		filePath := filepath.Join(rootPath, file.Name())
		if !utils.ExistsFile(filePath) {
			return
		}

		if err := os.Remove(filePath); err != nil {
			panic(err)
		}
	}
}

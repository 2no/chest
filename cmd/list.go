package cmd

import (
	"../utils"
	"fmt"
	"io/ioutil"
)

func List() {
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

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

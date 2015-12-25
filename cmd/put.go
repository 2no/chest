package cmd

import (
	"bufio"
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/wakuworks/chest/utils"
	"os"
	"path/filepath"
)

func Put(files []string) {
	chestPath := utils.GetChestPath()
	gitIgnorePath := utils.GetGitIgnorePath()

	for _, file := range files {
		if !utils.ExistsFile(file) {
			fmt.Println("File not found: " + file)
			return
		}

		if !utils.ExistsFile(chestPath) {
			os.Mkdir(chestPath, 0755)
		}

		filePath := filepath.Join(chestPath, file)
		if utils.ExistsFile(filePath) {
			fmt.Println("File already exist in chest: " + file)
			return
		}

		os.Rename(file, filePath)

		gitignore := false
		if utils.ExistsFile(gitIgnorePath) {
			fp, err := os.Open(gitIgnorePath)
			if err != nil {
				panic(err)
			}

			defer fp.Close()
			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				if "/"+file == scanner.Text() {
					gitignore = true
				}
			}
		}

		if !gitignore && prompter.YN("Add "+file+" to gitignore on the project root", true) {
			fp, err := os.OpenFile(gitIgnorePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			defer fp.Close()
			if _, err = fp.WriteString("/" + file + "\n"); err != nil {
				panic(err)
			}
		}
	}
}

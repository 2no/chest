package utils

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"os"
	"regexp"
)

func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func PrintFileDiff(file1 string, file2 string) {
	src1, _ := ioutil.ReadFile(file1)
	src2, _ := ioutil.ReadFile(file2)
	re := regexp.MustCompile("\r?\n")
	for _, diff := range lineDiff(string(src1), string(src2)) {
		if diff.Type == 0 {
			fmt.Printf(diff.Text)
		} else {
			for _, line := range re.Split(diff.Text, -1) {
				if diff.Type == -1 {
					fmt.Print("\x1b[31m-")
				} else {
					fmt.Print("\x1b[32m+")
				}
				fmt.Printf("%s\x1b[0m\n", line)
			}
		}
	}
}

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	return result
}

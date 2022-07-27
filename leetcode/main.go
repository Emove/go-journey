package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("hello leetcode")
	//fmt.Println(isNumber("23123"))
	renamePackages()
}

func renamePackages() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd = wd + "/code"
	_ = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() && err == nil {
			splits := strings.Split(info.Name(), "_")
			lastIndex := len(splits) - 1
			if !isNumber(splits[0]) && isNumber(splits[lastIndex]) {
				lastText := splits[lastIndex]
				copy(splits[1:], splits)
				splits[0] = lastText
				filename := strings.Join(splits, "_")
				newPath := path[:len(path)-len(filename)] + filename
				if err = os.Rename(path, newPath); err != nil {
					log.Fatal(err)
					return err
				}
				fmt.Printf("%s ----> %s\n", path, newPath)
			}
		}
		return nil
	})
}

func isNumber(str string) bool {
	for _, v := range str {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

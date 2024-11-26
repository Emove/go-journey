package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walk(path string) error {
	return filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(subpath)
		}
		return nil
	})
}

func Parent(path string) string {
	return filepath.Dir(filepath.Dir(path))
}

func CurrentDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

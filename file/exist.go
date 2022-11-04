package file

import "os"

func State(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

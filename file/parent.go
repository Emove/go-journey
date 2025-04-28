package file

import "path/filepath"

func GetDir(name string) string {
	return filepath.Dir(name)
}

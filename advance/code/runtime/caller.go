package runtime

import (
	"runtime"
)

func GetCaller(depth int) string {
	_, file, _, _ := runtime.Caller(depth)
	return file
}

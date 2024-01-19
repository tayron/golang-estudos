package util

import (
	"runtime"
	"strings"
)

func GetRuntimeFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	fullPathAtFunctionName := runtime.FuncForPC(pc).Name()
	listFullPathAtFunctionName := strings.Split(fullPathAtFunctionName, "/")
	return listFullPathAtFunctionName[len(listFullPathAtFunctionName)-1]
}

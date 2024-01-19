package util

import (
	"strings"
	"testing"
)

func TestGetRuntimeFunctionName(t *testing.T) {
	expectedFunctionName := "util.TestGetRuntimeFunctionName"
	fullPathAtFunctionName := GetRuntimeFunctionName()

	listFullPathAtFunctionName := strings.Split(fullPathAtFunctionName, "/")
	functionName := listFullPathAtFunctionName[len(listFullPathAtFunctionName)-1]

	if functionName != expectedFunctionName {
		t.Errorf("Expected function name: %s, but got: %s", expectedFunctionName, functionName)
	}
}

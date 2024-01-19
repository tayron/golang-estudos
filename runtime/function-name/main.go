package main

import (
	"fmt"

	"github.com/tayron/golang-estudos/runtime/function-name/util"
)

func main() {
	functionName := util.GetRuntimeFunctionName()
	fmt.Printf("Nome da função: %s\n", functionName)
}

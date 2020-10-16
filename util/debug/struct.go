package debug

import (
	"encoding/json"
	"fmt"
	"log"
)

// DebugarStruct -
func DebugarStruct(dado interface{}) {
	empJSON, err := json.MarshalIndent(dado, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", string(empJSON))
}

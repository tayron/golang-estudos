package method

import (
	"fmt"
	"net/http"
)

func HelloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World\n")
}

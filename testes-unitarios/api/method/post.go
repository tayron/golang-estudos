package method

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string
}

func HelloYou(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var p Person
	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", p.Name)

}

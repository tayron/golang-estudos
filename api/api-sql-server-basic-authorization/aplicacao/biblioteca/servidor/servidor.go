package servidor

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Startar -
func Startar(router http.Handler) {
	porta := os.Getenv("PORTA_EXECUCAO_APLICACAO")
	fmt.Println("Server executing at port " + porta)

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", porta), router)

	if err != nil {
		log.Fatal(err)
	}
}

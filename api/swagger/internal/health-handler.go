package internal

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/health api HealthHandler
//
// Verifica se a API está online e funcionando corretamente.
//
// Esta operação retorna um status HTTP 200 OK se a API estiver funcionando, indicando que ela está ativa e disponível.
//
// ---
// security:
// - apiKey: []
// responses:
//
//	'200':
//	  description: A API está funcionando corretamente
//	  content:
//	    text/plain:
//	      schema:
//	        type: string
//	        example: "OK"
//	'500':
//	  description: Erro interno do servidor
//	  content:
//	    text/plain:
//	      schema:
//	        type: string
//	        example: "Erro ao verificar a saúde da API"
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health endpoint foi chamado")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

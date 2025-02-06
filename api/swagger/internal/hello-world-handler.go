package internal

import (
	"net/http"
)

// swagger:route GET /api/hello api HelloWorldHandler
//
// Verifica se a API está online e funcionando corretamente.
//
// Esta operação retorna uma mensagem Hello world com status HTTP 200
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
//	        example: "Hello world"
//	'500':
//	  description: Erro interno do servidor
//	  content:
//	    text/plain:
//	      schema:
//	        type: string
//	        example: "Erro ao chamar hello world"
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

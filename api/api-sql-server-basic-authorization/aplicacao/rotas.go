package aplicacao

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tayron/golang-estudos/api/api-sql-server-basic-authorization/aplicacao/handlers"
)

// ConfigurarRotas -
func ConfigurarRotas() http.Handler {
	router := chi.NewRouter()
	router.Get("/api/v1/clientes", handlers.GetClientes)
	router.Get("/api/v1/cliente", handlers.GetClientesPorCPF)
	return router
}

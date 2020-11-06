package aplicacao

import (
	"net/http"

	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/handlers"
	"github.com/go-chi/chi"
)

// ConfigurarRotas -
func ConfigurarRotas() http.Handler {
	router := chi.NewRouter()
	router.Get("/api/v1/clientes", handlers.GetClientes)
	router.Get("/api/v1/cliente", handlers.GetClientesPorCPF)
	return router
}

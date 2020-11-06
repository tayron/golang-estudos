package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/util"
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/modelo"
)

// GetClientes - Retorna lista de clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {

	err := util.ValidarPermissaoAcesso(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	pagina := r.FormValue("pagina")
	numeroRegistros := r.FormValue("limit")
	listaClientes := modelo.ObterTodosClientes(pagina, numeroRegistros)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listaClientes)
}

// GetClientesPorCPF - Busca cliente por CPF
func GetClientesPorCPF(w http.ResponseWriter, r *http.Request) {

	err := util.ValidarPermissaoAcesso(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	cpf := r.FormValue("cpf")
	cliente := modelo.ObterClientePorCPF(cpf)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

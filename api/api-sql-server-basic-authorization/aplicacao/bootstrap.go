package aplicacao

import (
	"path/filepath"

	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/servidor"
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/util"

	"github.com/joho/godotenv"
)

// Inicializar - Método que starta algumas configurações do sistema
func Inicializar() {
	carregarConfiguracoesDoSistema()
	rotas := ConfigurarRotas()
	servidor.Startar(rotas)
}

func carregarConfiguracoesDoSistema() {
	caminhoAplicacao := util.ObterCaminhoDiretorioAplicacao()
	environmentPath := filepath.Join(caminhoAplicacao, ".env")
	err := godotenv.Load(environmentPath)
	if err != nil {
		panic(err)
	}
}

package dias

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// ObterDataInicio -
func ObterDataInicio() (string, error) {
	numeroDias, err := strconv.Atoi(os.Getenv("NUMERO_DIAS"))

	if err != nil {
		return "", errors.New("Não foi possível obter número de dias: " + err.Error())
	}

	numeroDias = numeroDias * -1

	dataAtual := time.Now()

	if numeroDias == 0 {
		return "2000-01-01 00:00:00", nil
	}

	if numeroDias == -1 {
		return dataAtual.Format("2006-01-02") + " 00:00:00", nil
	}

	numeroDias++

	dataInicial := dataAtual.AddDate(0, 0, numeroDias)
	dataInicio := dataInicial.Format("2006-01-02") + " 23:59:59"

	return dataInicio, nil
}

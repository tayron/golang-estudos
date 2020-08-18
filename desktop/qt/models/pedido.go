package models

import (
	"github.com/tayron/go-lang-estudos/desktop/qt/form/database"
)

type Pedido struct {
	ID        int
	Valor     int
	ValorPago int
}

func AtualizarValorPedido(pedidoMagento string, valor string) {
	db := database.ObterConexao()

	stmt, err := db.Prepare("update orders set amount_paid = ? where magento = ?")

	if err != nil {
		panic(err)
	}

	stmt.Exec(valor, pedidoMagento)
}

func ObterPedido(pedidoMagento string) Pedido {
	db := database.ObterConexao()

	rows, _ := db.Query("select id, amount, amount_paid from orders where magento = ?", pedidoMagento)
	defer rows.Close()

	var p Pedido
	for rows.Next() {
		rows.Scan(&p.ID, &p.Valor, &p.ValorPago)
		return p
	}

	return p
}

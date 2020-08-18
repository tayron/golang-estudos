package main

import (
	"errors"

	"github.com/leekchan/accounting"
	"github.com/tayron/go-lang-estudos/desktop/qt/models"
)

const (
	mensagemAtualizacaoPedido string  = "Pedido de id %s teve seu valor pago atualizado para: R$ %s"
	multiplicador             float64 = 0.01
)

// ObterValorPedido - Obtem valor do pedido Magento e o valor pago pelo cliente
func ObterValorPedido(pedidoMagento string) (valorPedido string, valorPago string, erro error) {
	pedido := models.ObterPedido(pedidoMagento)

	if pedido.ID == 0 {
		return "", "", errors.New("Pedido não encontrado")
	}

	var valorPedidoFloat float64 = (float64(pedido.Valor) * multiplicador)
	var valorPagoFloat float64 = (float64(pedido.ValorPago) * multiplicador)

	ac := accounting.Accounting{Symbol: "", Precision: 2, Thousand: ".", Decimal: ","}

	valorPedido = ac.FormatMoney(valorPedidoFloat)
	valorPago = ac.FormatMoney(valorPagoFloat)

	return
}

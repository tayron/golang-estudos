package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tayron/go-lang-estudos/desktop/qt/form/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leekchan/accounting"
	"github.com/therecipe/qt/widgets"
)

var multiplicador float64 = 0.01

func main() {
	godotenv.Load()
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Sistema Teste")
	window.SetMinimumSize2(300, 200)

	layout := widgets.NewQVBoxLayout()

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	labelPedidoMagento := widgets.NewQLabel2("Pedido", nil, 0)
	inputPedidoMagento := widgets.NewQLineEdit(nil)

	labelValorPedido := widgets.NewQLabel2("Valor Pedido", nil, 0)
	inputValorPedido := widgets.NewQLineEdit(nil)
	inputValorPedido.SetReadOnly(true)

	labelValorPago := widgets.NewQLabel2("Valor Pago", nil, 0)
	inputValorPago := widgets.NewQLineEdit(nil)

	layout.AddWidget(labelPedidoMagento, 0, 0)
	layout.AddWidget(inputPedidoMagento, 0, 0)
	layout.AddWidget(labelValorPedido, 0, 0)
	layout.AddWidget(inputValorPedido, 0, 0)
	layout.AddWidget(labelValorPago, 0, 0)
	layout.AddWidget(inputValorPago, 0, 0)

	buttonBuscarPedido := widgets.NewQPushButton2("Buscar", nil)
	buttonBuscarPedido.ConnectClicked(func(checked bool) {

		var pedidoMagento string = inputPedidoMagento.Text()
		pedido := models.ObterPedido(pedidoMagento)

		var valorPedidoFloat float64 = (float64(pedido.Valor) * multiplicador)
		var valorPagoFloat float64 = (float64(pedido.ValorPago) * multiplicador)

		ac := accounting.Accounting{Symbol: "", Precision: 2, Thousand: ".", Decimal: ","}
		valorPagoFormatado := ac.FormatMoney(valorPagoFloat)
		valorPedidoFormatado := ac.FormatMoney(valorPedidoFloat)

		inputValorPedido.SetText(valorPedidoFormatado)
		inputValorPago.SetText(valorPagoFormatado)
	})

	buttonAtualizarPedido := widgets.NewQPushButton2("Alterar", nil)

	buttonAtualizarPedido.ConnectClicked(func(checked bool) {

		var pedidoMagento string = inputPedidoMagento.Text()

		r := strings.NewReplacer(",", "", ".", "")
		var valor string = r.Replace(inputValorPago.Text())

		models.AtualizarValorPedido(pedidoMagento, valor)
		pedido := models.ObterPedido(pedidoMagento)

		var valorPedidoFloat float64 = (float64(pedido.Valor) * multiplicador)
		var valorPagoFloat float64 = (float64(pedido.ValorPago) * multiplicador)

		ac := accounting.Accounting{Symbol: "", Precision: 2, Thousand: ".", Decimal: ","}
		valorPagoFormatado := ac.FormatMoney(valorPagoFloat)
		valorPedidoFormatado := ac.FormatMoney(valorPedidoFloat)

		inputValorPedido.SetText(valorPedidoFormatado)
		inputValorPago.SetText(valorPagoFormatado)

		var retorno string = fmt.Sprintf("Pedido de id %d teve seu valor pago atualizado para: R$ %s", pedido.ID, valorPagoFormatado)

		widgets.QMessageBox_Information(nil, "Atualizado", retorno, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	layout.AddWidget(buttonBuscarPedido, 0, 0)
	layout.AddWidget(buttonAtualizarPedido, 0, 0)
	window.SetCentralWidget(widget)

	window.Show()
	widgets.QApplication_Exec()
}

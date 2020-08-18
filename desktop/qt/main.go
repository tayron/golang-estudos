package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tayron/go-lang-estudos/desktop/qt/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/therecipe/qt/widgets"
)

func init() {
	godotenv.Load()
	widgets.NewQApplication(len(os.Args), os.Args)
}

func main() {
	janela := GerarJanelaAplicacao()
	widget := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQVBoxLayout()
	widget.SetLayout(layout)

	labelPedidoMagento := widgets.NewQLabel2("Pedido Magento", nil, 0)
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
	buttonAtualizarPedido := widgets.NewQPushButton2("Alterar", nil)

	buttonBuscarPedido.ConnectClicked(func(checked bool) {

		var pedidoMagento string = inputPedidoMagento.Text()
		valorPedido, valorPago, erro := ObterValorPedido(pedidoMagento)

		if erro != nil {
			SetMensagemSucesso("Pedido não encontrado", erro.Error())
			return
		}

		SetTamanhoJanela(janela, 300, 270)

		inputValorPedido.SetText(valorPedido)
		inputValorPago.SetText(valorPago)
		layout.AddWidget(buttonAtualizarPedido, 0, 0)
	})

	buttonAtualizarPedido.ConnectClicked(func(checked bool) {

		var pedidoMagento string = inputPedidoMagento.Text()
		var valorPago string = RemoveMascaraValorMonetario(inputValorPago.Text())

		models.AtualizarValorPedido(pedidoMagento, valorPago)

		valorPedido, valorPago, erro := ObterValorPedido(pedidoMagento)

		if erro != nil {
			widgets.QMessageBox_Information(nil, "Erro", erro.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		inputValorPedido.SetText(valorPedido)
		inputValorPago.SetText(valorPago)

		var mensagemSucesso string = fmt.Sprintf(mensagemAtualizacaoPedido, pedidoMagento, valorPago)
		SetMensagemSucesso("Sucesso", mensagemSucesso)
	})

	layout.AddWidget(buttonBuscarPedido, 0, 0)

	janela.SetCentralWidget(widget)
	janela.Show()

	widgets.QApplication_Exec()
}

func GerarJanelaAplicacao() *widgets.QMainWindow {
	janela := widgets.NewQMainWindow(nil, 0)
	janela.SetWindowTitle("Pedido do Painel do Franqueado")
	SetTamanhoJanela(janela, 300, 200)

	return janela
}

func SetTamanhoJanela(janela *widgets.QMainWindow, largura int, altura int) {
	janela.SetMinimumSize2(largura, altura)
	janela.SetMaximumSize2(largura, largura)
}

func SetMensagemSucesso(tipo string, mensagem string) {
	widgets.QMessageBox_Information(nil, tipo, mensagem, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func RemoveMascaraValorMonetario(valorMonetario string) string {
	r := strings.NewReplacer(",", "", ".", "")
	return r.Replace(valorMonetario)
}

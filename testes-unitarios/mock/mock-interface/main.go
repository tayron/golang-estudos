package mockinterface

import "fmt"

type Notificador interface {
	Enviar(mensagem string) error
}

type ServicoBoasVindas struct {
	notificador Notificador
}

func NewServicoBoasVindas(notificador Notificador) ServicoBoasVindas {
	return ServicoBoasVindas{notificador: notificador}
}

func (s ServicoBoasVindas) BoasVindas(nome string) error {
	mensagem := fmt.Sprintf("Bem-vindo, %s!", nome)
	return s.notificador.Enviar(mensagem)
}

type NotificadorConsole struct{}

func (NotificadorConsole) Enviar(mensagem string) error {
	fmt.Println(mensagem)
	return nil
}

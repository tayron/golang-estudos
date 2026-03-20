package mockinterface

import "testing"

type MockNotificador struct {
	mensagemRecebida string
	foiChamado       bool
	erroRetornado    error
}

func (m *MockNotificador) Enviar(mensagem string) error {
	m.foiChamado = true
	m.mensagemRecebida = mensagem
	return m.erroRetornado
}

func TestServicoBoasVindasComMock(t *testing.T) {
	mock := &MockNotificador{}
	servico := NewServicoBoasVindas(mock)

	err := servico.BoasVindas("Tayron")
	if err != nil {
		t.Fatalf("nao esperava erro, mas recebeu %v", err)
	}

	if !mock.foiChamado {
		t.Fatal("esperava que o mock fosse chamado")
	}

	esperado := "Bem-vindo, Tayron!"
	if mock.mensagemRecebida != esperado {
		t.Fatalf("mensagem inesperada. esperava %q, recebeu %q", esperado, mock.mensagemRecebida)
	}
}

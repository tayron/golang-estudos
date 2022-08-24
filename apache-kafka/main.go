package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func main() {
	topico := "processar-despesa"
	mensagem := "Hello World"
	adicionarMensagemNaFila(topico, mensagem)
	lerMensagemNaFila(topico)

}

func adicionarMensagemNaFila(topico string, mensagem string) {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:29092"),
		Topic: topico,
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(mensagem),
		Headers: []protocol.Header{
			{
				Key:   "session",
				Value: []byte("123"),
			},
		},
	})

	if err != nil {
		log.Fatal("Não foi possível escrever a mensagem: ", err)
	}
}

func lerMensagemNaFila(topico string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:29092"},
		GroupID:  "consumer",
		Topic:    topico,
		MinBytes: 0,
		MaxBytes: 10e6, //10MB
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Println("Sessão 123 encontrada!")
			}
		}

		if err != nil {
			log.Fatal("Não foi possível receber a mensagem: ", err)
			reader.Close()
		}

		fmt.Println("Mensagem recebida: ", string(message.Value))
	}

	reader.Close()
}

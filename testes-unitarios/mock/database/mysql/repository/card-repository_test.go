package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCardRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o mock do banco: %s", err)
	}
	defer db.Close()

	repo := NewCardRepository(db)

	card := Card{Number: 1234}

	mock.ExpectPrepare("insert into cards values \\(\\?\\)").
		ExpectExec().
		WithArgs(card.Number).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(card)
	assert.NoError(t, err)

	// Assegurando que todas as expectativas do mock foram atendidas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectativas não foram atendidas: %s", err)
	}
}

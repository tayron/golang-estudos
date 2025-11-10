package repository

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func NewMockDB() (*sql.DB, error) {
	// Cria um banco de dados SQLite in-memory para testes
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Cria a tabela necessária para os testes
	_, err = db.Exec("CREATE TABLE cards (number INTEGER)")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateCard(t *testing.T) {
	db, err := NewMockDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewCardRepository(db)

	card := Card{Number: 12345}
	err = repo.Create(card)
	assert.NoError(t, err)

	// Verifica se o card foi inserido corretamente
	var number int
	err = db.QueryRow("SELECT number FROM cards WHERE number = ?", card.Number).Scan(&number)
	assert.NoError(t, err)
	assert.Equal(t, card.Number, number)
}

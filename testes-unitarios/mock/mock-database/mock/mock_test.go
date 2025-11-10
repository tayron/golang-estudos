package mock

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// 1️⃣ Cria o mock e o "db fake"
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, "erro ao criar o mock do banco de dados:", err.Error())
	}
	defer db.Close()

	// 2️⃣ Define o comportamento esperado:
	// Quando uma query "SELECT name FROM users WHERE id = ?" for executada com argumento 10,
	// o mock deve retornar uma linha com o valor "Tayron".
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Tayron")
	mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
		WithArgs(10).
		WillReturnRows(rows)

	// 3️⃣ Executa a query normalmente, como se fosse um banco real
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = ?", 10).Scan(&name)
	if err != nil {
		assert.FailNow(t, "erro ao executar a query no banco de dados:", err.Error())
	}

	fmt.Println("Resultado retornado pelo banco:", name)

	assert.Equal(t, "Tayron", name, "o nome retornado deve ser 'Tayron'")

	// 4️⃣ Verifica se o comportamento esperado realmente foi executado
	assert.NoError(t, mock.ExpectationsWereMet())
}

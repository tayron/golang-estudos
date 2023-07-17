package databasemock

import (
	"example.com/mock/crud-sqlite/interface/database"
	"exemple.com/crud-sqlite/entity"
)

type DatabaseMock struct {
	Users map[int]entity.User
}

func (m *DatabaseMock) Exec(query string, args ...interface{}) (database.Result, error) {
	// Implemente a execução simulada para o mock.
	return nil, nil
}

func (m *DatabaseMock) Query(query string, args ...interface{}) (*database.Rows, error) {
	// Implemente a consulta simulada para o mock.
	return nil, nil
}

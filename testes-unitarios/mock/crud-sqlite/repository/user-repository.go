package repository

import (
	"example.com/mock/crud-sqlite/entity"
	"example.com/mock/crud-sqlite/interface/database"
)

func CreateUser(db database.Database, name string, age int) (int64, error) {
	// Implemente a criação do usuário usando a interface DB.
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetUser(db database.Database, id int) (*entity.User, error) {
	// Implemente a recuperação do usuário usando a interface DB.
	rows, err := db.Query("SELECT id, name, age FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, nil
}

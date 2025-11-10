package repository

import "database/sql"

type Card struct {
	Number int
}

type CardRepositoryInterface interface {
	Create(card Card) error
}

type CardRepository struct {
	db *sql.DB
}

func NewCardRepository(db *sql.DB) CardRepositoryInterface {
	return CardRepository{db}
}

func (r CardRepository) Create(card Card) error {
	query := "insert into cards values (?)"
	smt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = smt.Exec(card.Number)
	if err != nil {
		return err
	}

	return nil
}

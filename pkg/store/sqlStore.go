package store

import (
	"database/sql"
	"repositoryapi/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Read(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM products WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.ID, &dentist.Name, &dentist.LastName, &dentist.Registration)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

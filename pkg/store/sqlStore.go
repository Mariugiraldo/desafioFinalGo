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
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.Registration)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStore) CreateDentist(dentist domain.Dentist) (domain.Dentist, error) {
	query := "INSERT INTO dentists (id, name, lastname, registration) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	stmt.Exec(dentist.Id, dentist.Name, dentist.LastName, dentist.Registration )

	return dentist, nil
}

func (s *sqlStore) UpdateDentist(dentist domain.Dentist) (domain.Dentist, error) {
	query := "UPDATE dentists SET name =?, lastname =?, registration =? WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	stmt.Exec(dentist.Name, dentist.LastName, dentist.Registration, dentist.Id)

	return dentist, nil
}

func (s *sqlStore) PatchDentist(dentist domain.Dentist) (domain.Dentist, error) {
	query := "UPDATE dentists SET name =? WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	stmt.Exec(dentist.Name, dentist.Id)

	return dentist, nil
}

func (s *sqlStore) Delete(id int ) {
	query := "DELETE FROM dentists WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return 
	}
	stmt.Exec(id)

	return 
}
package store

import (
	"database/sql"
	"log"
	"repositoryapi/internal/domain"
)

func NewSqlStoreShift(db *sql.DB) ShiftStoreInterface {
	return &sqlStore{db}

}

func (s *sqlStore) ReadAllShift() ([]domain.Shift, error) {
	query := "SELECT * FROM shifts;"
	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.Shift{}, err
	}
	defer rows.Close()
	return []domain.Shift{}, nil

}

func (s *sqlStore) ReadShift(id int) (domain.Shift, error) {
	var shift domain.Shift
	query := "SELECT * FROM shifts WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&shift.ID, &shift.PatientID, &shift.DentistID, &shift.DischargeDate, &shift.Description)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil

}

func (s *sqlStore) CreateShift(shift domain.Shift) (domain.Shift, error) {
	query := "INSERT INTO shifts (patient_id, dentist_id, DischargeDate, description) VALUES (?,?,?,?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(shift.PatientID, shift.DentistID, shift.DischargeDate, shift.Description)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Shift{}, err
	}

	defer stmt.Close()
	return shift, nil

}

func (s *sqlStore) UpdateShift(shift domain.Shift) (domain.Shift, error) {
	query := "UPDATE shifts SET patient_id = ?, dentist_id = ?, DischargeDate = ?, description = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(shift.PatientID, shift.DentistID, shift.DischargeDate, shift.Description, shift.ID)

	return shift, nil
}

func (s *sqlStore) PatchShift(shift domain.Shift) (domain.Shift, error) {
	query := "UPDATE shifts SET description = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(shift.Description, shift.ID)

	return shift, nil

}

func (s *sqlStore) DeleteShift(id int) error {
	query := "DELETE FROM shifts WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	stmt.Exec(id)
	return nil
}

func (s *sqlStore) ExistShift(id int) bool {
	query := "SELECT * FROM shifts WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	var shift domain.Shift
	err := row.Scan(&shift.ID, &shift.PatientID, &shift.DentistID, &shift.DischargeDate, &shift.Description)
	if err != nil {
		return false
	}
	return true
}

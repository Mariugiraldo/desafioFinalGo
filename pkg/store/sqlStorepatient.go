package store

import (
	"database/sql"
	"log"
	"repositoryapi/internal/domain"
	"time"
)

func NewSqlStorePatient(db *sql.DB) PatientStoreInterface {
	return &sqlStore{db}

}

func (s *sqlStore) ReadAllPatient() ([]domain.Patient, error) {
	query := "SELECT * FROM patients;"
	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.Patient{}, err
	}
	defer rows.Close()
	return []domain.Patient{}, nil
}

func (s *sqlStore) ReadPatient(id int) (domain.Patient, error) {
	var patient domain.Patient
	var dischargeDateStr string // temporary variable to hold the discharge date as a string
	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.ID, &patient.Name, &patient.LastName, &patient.Home, &patient.DNI, &dischargeDateStr) // scan the discharge date into the temporary string variable
	if err != nil {
		return domain.Patient{}, err
	}

	// Convert dischargeDateStr from string to time.Time
	dischargeDate, err := time.Parse("2006-01-02", dischargeDateStr)
	if err != nil {
		return domain.Patient{}, err
	}
	patient.DischargeDate = dischargeDate // now you can assign the time.Time value to DischargeDate

	return patient, nil
}

func (s *sqlStore) CreatePatient(pat domain.Patient) (domain.Patient, error) {
	query := "INSERT INTO patients (name, lastName, Home, DNI, DischargeDate) VALUES (?,?,?,?,?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(pat.Name, pat.LastName, pat.Home, pat.DNI, pat.DischargeDate)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	defer stmt.Close()
	return pat, nil
}

func (s *sqlStore) UpdatePatient(pat domain.Patient) (domain.Patient, error) {
	query := "UPDATE patients SET name = ?, lastName = ?, home = ?, DNI = ?, dischargeDate = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(pat.Name, pat.LastName, pat.Home, pat.DNI, pat.DischargeDate, pat.ID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	defer stmt.Close()
	return pat, nil
}

func (s *sqlStore) PatchPatient(patient domain.Patient) (domain.Patient, error) {
	query := "UPDATE patients SET name =? WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Patient{}, err
	}
	stmt.Exec(patient.Name, patient.ID)

	return patient, nil
}

func (s *sqlStore) DeletePatient(id int) error {
	query := "DELETE FROM patients WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func (s *sqlStore) ExistPatient(id int) bool {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM patients WHERE id = ?;", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

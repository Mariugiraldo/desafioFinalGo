package store

import (
	"database/sql"
	"fmt"
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
	stmt.Exec(dentist.Id, dentist.Name, dentist.LastName, dentist.Registration)

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

func (s *sqlStore) Delete(id int) {
	query := "DELETE FROM dentists WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	stmt.Exec(id)

	return
}

func (s *sqlStore) ReadShift(id int) (domain.Shift, error) {
	var shift domain.Shift
	query := "SELECT * FROM shifts WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&shift.ID, &shift.PatientID, &shift.DentistID, &shift.DischargeDate, &shift.Description)
	fmt.Println("shift details: ", shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (s *sqlStore) CreateShift(shift domain.Shift) (domain.Shift, error) {
	query := "INSERT INTO shifts (id, patient_id, dentist_id, dischargedate, description) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Shift{}, err
	}

	result, err := stmt.Exec(shift.ID, shift.PatientID, shift.DentistID, shift.DischargeDate, shift.Description)
	if err != nil {
		fmt.Println(err)
		return domain.Shift{}, err
	}
	result.RowsAffected()
	return shift, nil
}

func (s *sqlStore) UpdateShift(shift domain.Shift) (domain.Shift, error) {
	query := "UPDATE shifts SET patient_id =?, dentist_id =?, dischargedate =?, description =? WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return domain.Shift{}, err
	}
	result, err := stmt.Exec(shift.PatientID, shift.DentistID, shift.DischargeDate, shift.Description, shift.ID)
	if err != nil {
		fmt.Println(err)
		return domain.Shift{}, err

	}
	result.RowsAffected()

	return shift, nil
}

func (s *sqlStore) DeleteShift(id int) {
	query := "DELETE FROM shifts WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	stmt.Exec(id)

	return
}

func (s *sqlStore) PatchShift(shift domain.Shift) (domain.Shift, error) {
	query := "UPDATE shifts SET description =? WHERE id =?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Shift{}, err
	}
	stmt.Exec(shift.Description, shift.ID)

	return shift, nil
}

func (s *sqlStore) CreateShiftByDNIAndRegistration(dni string, registration string, shift domain.Shift) (domain.Shift, error) {
	query := "SELECT * FROM patients WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	var patient domain.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.LastName, &patient.Home, &patient.DNI, &patient.DischargeDate)
	fmt.Println("patient details: ", patient)
	if err != nil {
		return domain.Shift{}, err
	}
	query = "SELECT * FROM dentists WHERE registration = ?;"
	row = s.db.QueryRow(query, registration)
	var dentist domain.Dentist
	err = row.Scan(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.Registration)
	fmt.Println("dentist details: ", dentist)
	if err != nil {
		return domain.Shift{}, err
	}
	query = "INSERT INTO shifts (patient_id, dentist_id, dischargedate, description) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Shift{}, err
	}
	result, err := stmt.Exec(patient.ID, dentist.Id, shift.DischargeDate, shift.Description)
	if err != nil {
		return domain.Shift{}, err
	}
	result.RowsAffected()
	return shift, nil
}


func (s *sqlStore) ReadShiftByDNI(dni string) (domain.Shift, error) {
	query := "SELECT * FROM patients WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	var patient domain.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.LastName, &patient.Home, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return domain.Shift{}, err
	}

	query = "SELECT * FROM shifts WHERE patient_id = ?;"
	row = s.db.QueryRow(query, patient.ID)
	var shift domain.Shift
	err = row.Scan(&shift.ID, &shift.PatientID, &shift.DentistID, &shift.DischargeDate, &shift.Description)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

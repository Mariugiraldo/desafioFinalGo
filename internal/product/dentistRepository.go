package product

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type DentistRepository interface {
	GetByID(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist)(domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist)(domain.Dentist, error)
	PatchDentist(dentist domain.Dentist)(domain.Dentist, error)
	DeleteDentist(id int) 
	
}

type dentistRepository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) DentistRepository {
	return &dentistRepository{storage}
}

func (r *dentistRepository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")

	}
	return dentist, nil
}

func (r *dentistRepository ) CreateDentist(dentist domain.Dentist)(domain.Dentist, error){
	
	return r.storage.CreateDentist(dentist)

}

func (r *dentistRepository ) UpdateDentist(dentist domain.Dentist)(domain.Dentist, error){
	
	return r.storage.UpdateDentist(dentist)

}


func (r *dentistRepository ) PatchDentist(dentist domain.Dentist)(domain.Dentist, error){
	
	return r.storage.PatchDentist(dentist)

}

func (r *dentistRepository ) DeleteDentist(id int){
	r.storage.Delete(id)
	return

}


package store

import "repositoryapi/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Dentist, error)
}

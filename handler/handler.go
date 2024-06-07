package handler

import "github.com/MatiasCeb/api-go/tree/main/model/model"

// Storage .
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}

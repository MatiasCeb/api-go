package model

import "errors"

var (
	//ErrCantBeNil: El ingreso no puede ser nulo
	ErrCantBeNil             = errors.New("El ingreso no puede ser nulo")
	ErrIDPersonDoesNotExists = errors.New("No existe una persona con ese ID")
)

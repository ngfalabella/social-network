package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID       int
	Nombre   string
	Email    string
	Password string
}

func ValidateUser(u User) error {
	if u.Email == "" {
		fmt.Println("Paso por el model")
		return errors.New("El campo email no puede estar vacio")
	}
	return nil
}

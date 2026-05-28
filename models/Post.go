package models

import "errors"

type Post struct {
	ID        int
	Autor     string
	Contenido string
}

func ValidatePost(p Post) error {
	if p.Contenido == "" {
		return errors.New("El campo de contenido no puede estar vacio")
	}
	if len(p.Contenido) > 40 {
		return errors.New("El post no puede tener mas de 40 caracteres")
	}
	return nil
}

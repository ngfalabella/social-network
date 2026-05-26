package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func main() {
	fmt.Println("Backend iniciado")

	firstUser := User{
		ID:       1,
		Username: "n_test",
		Email:    "n_test@gmail.com",
		Password: "clave_0123",
	}

	fmt.Println("Usuario creado con exito")

	fmt.Printf("ID USER : %d ,\n USERNAME: %s ,\n EMAIL: %s , \n PASSWORD: %s", firstUser.ID, firstUser.Username, firstUser.Email, firstUser.Password)

}

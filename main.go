package main

import (
	"fmt"

	"github.com/ngfalabella/social-network/models"
)

func main() {
	fmt.Println("Inicio de programa")

	usuarioUno := models.User{ID: 1, Nombre: "Juan", Email: "Juan@gmail.com", Password: "123"}
	usuarioDos := models.User{ID: 2, Nombre: "Pedro", Email: "Pedro@gmail.com", Password: "456"}

	var listaDeUsuarios []models.User

	listaDeUsuarios = append(listaDeUsuarios, usuarioUno, usuarioDos)

	fmt.Printf("Usuarios Cargados hasta el momento : %d \n", len(listaDeUsuarios))

	fmt.Println(listaDeUsuarios)

	fmt.Println("Cargando el proximo usuario")

	usuarioTres := models.User{ID: 3, Nombre: "Marcos", Email: "Marcos@gmail.com", Password: "789"}

	err := models.ValidateUser(usuarioTres)

	if err != nil {
		fmt.Println("Error encontrado ", err)
		return
	}

	listaDeUsuarios = append(listaDeUsuarios, usuarioTres)

	fmt.Println("Usuario cargado con exito")

	PrintUserReport(listaDeUsuarios)

	var feed []models.Post

	publicacionUno := models.Post{ID: 1, Autor: "Marcos", Contenido: "Programando en Go desde cero!"}
	publicacionDos := models.Post{ID: 2, Autor: "Marcos", Contenido: "Segundo post en Go"}
	publicacionTres := models.Post{ID: 3, Autor: "Marcos", Contenido: "Fin del dia!"}

	errDos := models.ValidatePost(publicacionUno)
	errTres := models.ValidatePost(publicacionDos)
	errCuatro := models.ValidatePost(publicacionTres)

	if errDos != nil {
		fmt.Println("Nuevo error : ", errDos)
	}
	if errTres != nil {
		fmt.Println("Nuevo error : ", errTres)
	}
	if errCuatro != nil {
		fmt.Println("Nuevo error : ", errCuatro)
	}

	feed = append(feed, publicacionUno, publicacionDos, publicacionTres)

	fmt.Println(feed)

}

func PrintUserReport(lista []models.User) {
	for _, elemento := range lista {
		fmt.Println("--- Perfil de Usuario ---")
		fmt.Printf("ID : %d \n", elemento.ID)
		fmt.Printf("Nombre : %s \n", elemento.Nombre)
		fmt.Printf("Email : %s \n", elemento.Email)
		fmt.Println("-------------------------")
	}
}

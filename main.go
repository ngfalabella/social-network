package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ngfalabella/social-network/models"
)

const PORT = 7070

var ListadoDePost = []models.Post{}

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no validado", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListadoDePost)
}

func CrearPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no validado", http.StatusMethodNotAllowed)
		return
	}
	var postCreado models.Post

	err := json.NewDecoder(r.Body).Decode(&postCreado)
	if err != nil {
		http.Error(w, "Error al leer datos", http.StatusBadRequest)
		return
	}

	ListadoDePost = append(ListadoDePost, postCreado)

	w.WriteHeader(http.StatusCreated)
	fmt.Println("Post creado ")
}

func init() {

	postUno := models.Post{ID: 222, Autor: "Adrian", Contenido: "Me fui de vacaciones", Likes: 77, IpCreador: "199.199.180.1"}

	ListadoDePost = append(ListadoDePost, postUno)
}

func main() {
	addr := fmt.Sprintf(":%d", PORT)

	http.HandleFunc("/api/posts/obtener", HandlerPost)
	http.HandleFunc("/api/posts/crear", CrearPostHandler)

	fmt.Printf("Red Social activa en http://localhost:%d", PORT)

	http.ListenAndServe(addr, nil)
}

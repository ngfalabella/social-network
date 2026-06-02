package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ngfalabella/social-network/models"
)

type Product struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Precio int    `json:"precio"`
	Costo  int    `json:"-"`
}

const PORT int = 3030

var ListadoDeProductos = []Product{}
var ListadoDeComentarios = []models.Coment{}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Bienvenido a la Tienda de la Red Social")
	}
	fmt.Fprint(w, "No se encontro ruta")
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListadoDeProductos)
}

func HandleComents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListadoDeComentarios)
}

func init() {
	productoUno := Product{ID: 1, Nombre: "Celular", Precio: 1000000, Costo: 750000}
	productoDos := Product{ID: 2, Nombre: "Notebook", Precio: 1250000, Costo: 800000}
	comentarioUno := models.Coment{ID: 123, Texto: "Muy buena publicacion", IPUsuario: "UD-789"}

	ListadoDeProductos = append(ListadoDeProductos, productoUno, productoDos)
	ListadoDeComentarios = append(ListadoDeComentarios, comentarioUno)

}

func main() {
	addr := fmt.Sprintf(":%d", PORT)

	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/api/products", HandleProducts)
	http.HandleFunc("/api/coments", HandleComents)

	fmt.Println("Servidor corriendo en Puerto : ", PORT)
	http.ListenAndServe(addr, nil)

}

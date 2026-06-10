package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Importamos tus paquetes internos usando el nombre de tu módulo
	"github.com/ngfalabella/social-network/internal/database"
	"github.com/ngfalabella/social-network/internal/handlers"

	
	"github.com/joho/godotenv"
)

func main() {
	// 1. 🔍 Cargamos el archivo .env antes de hacer cualquier otra cosa
	err := godotenv.Load("../../.env") // Usamos "../../" porque main.go ahora está metido adentro de cmd/server/
	if err != nil {
		log.Println("Aviso: No se encontró el archivo .env, se usarán las variables del sistema")
	}

	// 2. 🔌 Levantamos el puente con Postgres
	database.ConectarBaseDeDatos()
	defer database.DB.Close() // Nos aseguramos de cerrar la manguera al apagar la app

	// 3. 🗺️ Registramos al menos una ruta conectada a tus nuevos handlers independientes
	http.HandleFunc("/api/posts/obtener", handlers.HandlerPost)
	// http.HandleFunc("/api/posts/crear", handlers.HandleCreatePost) // La dejamos lista para cuando mudes este handler
	http.HandleFunc("/api/posts/crear", handlers.HandleCreatePost)

	// 4. 🚀 Leemos el puerto del .env y encendemos el servidor web
	port := os.Getenv("PORT")
	if port == "" {
		port = "7070" // Puerto de respaldo por si acaso
	}
	addr := fmt.Sprintf(":%s", port)

	fmt.Printf("🚀 Red Social activa y conectada en http://localhost:%s\n", port)
	
	// El servidor se queda escuchando acá las 24hs
	log.Fatal(http.ListenAndServe(addr, nil))
}
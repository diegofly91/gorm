package main

import (
	"gorm/src/modules"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	// Rutas
	mux := mux.NewRouter()
	modules.SetupModules(mux)
	mux.PathPrefix("/api").Handler(http.StripPrefix("/api", mux))
	// Servidor
	log.Fatal(http.ListenAndServe(":"+port, mux))

}

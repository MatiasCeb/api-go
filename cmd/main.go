package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatiasCeb/api-go/tree/main/model/authorization"
	"github.com/MatiasCeb/api-go/tree/main/model/handler"
	"github.com/MatiasCeb/api-go/tree/main/model/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Server iniciado")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error en el servidor: %v\n", err)
	}
}

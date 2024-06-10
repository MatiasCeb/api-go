package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatiasCeb/api-go/tree/main/model/handler"
	"github.com/MatiasCeb/api-go/tree/main/model/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Server iniciado")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error en el servidor: %v\n", err)
	}
}

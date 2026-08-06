package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP é a base da comunicação web, é um protocolo de comunicação")

	// ROTAS
	// URI - Identificador do recurso
	// método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola mundo"))
	})
	http.HandleFunc("/home/usuario", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola usuario"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}

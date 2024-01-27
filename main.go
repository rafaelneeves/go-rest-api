package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
	"go-rest-api/controllers"
)

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/listar", clienteController.BuscarCliente)
    r.HandleFunc("/inserir", clienteController.InserirCliente)
    r.HandleFunc("/atualizar", clienteController.AtualizarCliente)
    r.HandleFunc("/deletar", clienteController.DeletarCliente)

    log.Fatal(http.ListenAndServe(":8000", r))
}


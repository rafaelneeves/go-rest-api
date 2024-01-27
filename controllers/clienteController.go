package clienteController

import (
	"fmt"
	"net/http"
	"strconv"
	"go-rest-api/db"
	"encoding/json"
)

type Cliente struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	conexao := db.Conecta()
	linha, erro := conexao.Query("SELECT * FROM clientes")

	if erro != nil {
		fmt.Println("Erro ao executar a consulta:", erro)
		return
	}

	defer linha.Close()

	var clientes []Cliente

	for linha.Next() {
		var cliente Cliente

		if erro := linha.Scan(&cliente.Id, &cliente.Nome); erro != nil {
			fmt.Println("Erroo ao escanear a linha:", erro)
		}

		clientes = append(clientes, cliente)
	}	

	json, erro := json.Marshal(clientes)

	if erro != nil {
		fmt.Println("Erro ao converter para JSON:", erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.Write(json)
}

func InserirCliente(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")

	if nome == "" {
		fmt.Fprint(w, "Erro. Nome vazio ou inválido\n")
		return
	} 

	conexao := db.Conecta()
	inserir_produto, erro := conexao.Prepare("INSERT INTO clientes(nome) VALUES(?)")

	if erro != nil {
		fmt.Println("Erro ao inserir:", erro)
		return
	}

	inserir_produto.Exec(nome)
	defer conexao.Close()

	json, erro := json.Marshal(nome)
	
	if erro != nil {
		fmt.Println("Erro ao converter para JSON:", erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.Write(json)	
}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	idStr := r.URL.Query().Get("id")
	
	var resultado string

	if nome == "" || idStr == ""{
		fmt.Fprint(w, "Erro. Nome ou ID inválido\n")
		return
	} 
	
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Erro ao converter ID para inteiro:", err)
		return
	}

	conexao := db.Conecta()
	atualiza_produto, erro := conexao.Prepare("UPDATE clientes SET nome = ? WHERE id = ?")

	if erro != nil {
		fmt.Println("Erro ao atualizar:", erro)
		return
	}

	atualiza_produto.Exec(nome, id)

	defer conexao.Close()

	resultado = "Cliente id: " + idStr + " com o nome: " + nome + " Atualizado!"

	json, erro := json.Marshal(resultado)
	
	if erro != nil {
		fmt.Println("Erro ao converter para JSON:", erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.Write(json)	
}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == ""{
		fmt.Fprint(w, "Erro. ID inválido\n")
		return
	} 
	
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Erro ao converter ID para inteiro:", err)
		return
	}

	conexao := db.Conecta()
	deleta_produto, erro := conexao.Prepare("DELETE FROM clientes WHERE id = ?")
	
	if erro != nil {
		fmt.Println("Erro ao deletar:", erro)
	}
	
	deleta_produto.Exec(id)
	
	defer conexao.Close()

	var resultado string
	resultado = "Cliente id: " + idStr + " deletado"

	json, erro := json.Marshal(resultado)
	
	if erro != nil {
		fmt.Println("Erro ao converter para JSON:", erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.Write(json)	
}
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Cliente struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     int    `json:"idade"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
	Rua       string `json:"rua"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
	Cep       string `json:"cep"`
}

var data []Cliente = []Cliente{
	{
		Id:        1,
		Nome:      "João",
		Sobrenome: "Silva",
		Idade:     30,
		Email:     "1dIbK@example.com",
		Telefone:  "123456789",
		Rua:       "Rua A",
		Cidade:    "Cidade A",
		Estado:    "Estado A",
		Cep:       "12345-678",
	},
	{
		Id:        2,
		Nome:      "Maria",
		Sobrenome: "Santos",
		Idade:     25,
		Email:     "1dIbK@example.com",
		Telefone:  "123456789",
		Rua:       "Rua B",
		Cidade:    "Cidade B",
		Estado:    "Estado B",
		Cep:       "12345-678",
	},
	{
		Id:        3,
		Nome:      "Pedro",
		Sobrenome: "Almeida",
		Idade:     35,
		Email:     "1dIbK@example.com",
		Telefone:  "123456789",
		Rua:       "Rua C",
		Cidade:    "Cidade C",
		Estado:    "Estado C",
		Cep:       "12345-678",
	},
}
var maxId int = 0

// main inicia um servidor web na porta 8080 com endpoints para:
//
// - POST /client: cadastra um novo cliente com o corpo da requisi o em json
// - PUT /client/{id}: atualiza os dados de um cliente com o id especificado
//   com o corpo da requisi o em json
// - DELETE /client/{id}: remove um cliente com o id especificado
// - GET /client/{id}: retorna os dados de um cliente com o id especificado
// - GET /client: retorna os dados de todos os clientes
func main() {
	for _, v := range data {
		if v.Id > maxId {
			maxId = v.Id
		}
	}

	http.HandleFunc("POST /client", func(w http.ResponseWriter, r *http.Request) {
		cliente := Cliente{}
		err := json.NewDecoder(r.Body).Decode(&cliente)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente.Id = maxId + 1
		maxId++
		data = append(data, cliente)
		json.NewEncoder(w).Encode(cliente)
	})

	http.HandleFunc("PUT /client/{id}", func(w http.ResponseWriter, r *http.Request) {
		intId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		cliente := Cliente{}
		err = json.NewDecoder(r.Body).Decode(&cliente)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente.Id = intId
		for i, v := range data {
			if v.Id == intId {
				data[i] = cliente
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	http.HandleFunc("DELETE /client/{id}", func(w http.ResponseWriter, r *http.Request) {
		intId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, v := range data {
			if v.Id == intId {
				data = append(data[:i], data[i+1:]...)
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	http.HandleFunc("GET /client/{id}", func(w http.ResponseWriter, r *http.Request) {
		intId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, v := range data {
			if v.Id == intId {
				json.NewEncoder(w).Encode(v)
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	http.HandleFunc("GET /client", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)

}

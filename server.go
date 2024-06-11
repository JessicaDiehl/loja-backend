package main

import (
	"encoding/json"
	"net/http"

	"github.com/JessicaDiehl/loja-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProduto(w, r)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	err := registraProduto(produto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Erro{ErrorMessage: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query() //r representa toda a request, dentro da URL tem uma funcao query, retornando as queries passadas.

	nome := queryParam.Get("nome") //especificamos o parametro de query que tem que ser passado

	if nome != "" {
		produtosFiltradosPorNome := produtosPorNome(nome)
		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
	} else {
		Produtos := produtos
		json.NewEncoder(w).Encode(Produtos)
	}
}
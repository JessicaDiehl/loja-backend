package main

import (
	"errors"

	"github.com/JessicaDiehl/loja-backend/model"
)

var produtos []model.Produto = []model.Produto{}

func produtosPorNome(nome string) []model.Produto {

	resultados := []model.Produto{}

	for _, p := range produtos {
		if p.Nome == nome {
			resultados = append(resultados, p)
		}
	}
	return resultados
}

func registraProduto(produtoNovo model.Produto) error {

	for _, produto := range produtos {
		if produtoNovo.Id == produto.Id {
			return errors.New("produto com o ID ja cadastrado")
		}
	}
	produtos = append(produtos, produtoNovo)
	return nil
}
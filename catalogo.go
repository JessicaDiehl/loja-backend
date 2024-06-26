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

func atualizaProduto(produtoAtualizado model.Produto) error {
    for i, produto := range produtos {
        if produto.Id == produtoAtualizado.Id {
            produtos[i] = produtoAtualizado
            return nil
        }
    }
    return errors.New("produto não encontrado")
}

func removeProduto(id string) error {
    for i, produto := range produtos {
        if produto.Id == id {
            produtos = append(produtos[:i], produtos[i+1:]...)
            return nil
        }
    }
    return errors.New("produto não encontrado")
}
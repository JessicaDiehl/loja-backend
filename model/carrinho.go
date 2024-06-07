package model

type Carrinho struct {
	Id           string
	UserId       string
	ValorTotal   float64
	InfosProduto []InfosProduto
}

type InfosProduto struct {
	ProdutoId           string
	QuantidadeDeProduto int
}
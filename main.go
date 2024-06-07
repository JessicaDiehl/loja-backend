// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	fmt.Println("Esse é o catálogo da loja: ")

	nome := "Calça Karol" // substituir por um nome de produto

	produtosFiltrados := produtosPorNome(nome)

	fmt.Println(produtosFiltrados)

	StartServer()
}
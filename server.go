package main

import (
	"encoding/json"
	"net/http"

	"github.com/JessicaDiehl/loja-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)
	http.HandleFunc("/carrinhos", createCarrinho).Methods("POST")
  	http.HandleFunc("/carrinhos/{id}", getCarrinho).Methods("GET")
  	http.HandleFunc("/carrinhos/{id}", updateCarrinho).Methods("PUT")
  	http.HandleFunc("/carrinhos/{id}", deleteCarrinho).Methods("DELETE")
	http.HandleFunc("/usuarios", createUsuario).Methods("POST")
	http.HandleFunc("/wishlist", createItemWishlist).Methods("POST")
    http.HandleFunc("/wishlist", getItemsWishlist).Methods("GET")
    http.HandleFunc("/wishlist/{id}", deleteItemWishlist).Methods("DELETE")

    http.HandleFunc("/pagamento", createPagamento).Methods("POST")
    http.HandleFunc("/pagamento/{id}", getPagamento).Methods("GET")
        http.HandleFunc("/login", loginUsuario).Methods("POST")
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


func createCarrinho(w http.ResponseWriter, r *http.Request) {
    var novoCarrinho model.Carrinho
    _ = json.NewDecoder(r.Body).Decode(&novoCarrinho)
    novoCarrinho.Id = strconv.Itoa(rand.Intn(1000000)) // Gera um ID aleatório

    carrinhos[novoCarrinho.Id] = novoCarrinho

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(novoCarrinho)
}

func getCarrinho(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    carrinho, ok := carrinhos[id]

    if !ok {
        http.Error(w, "Carrinho não encontrado", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(carrinho)
}

func updateCarrinho(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var carrinhoAtualizado model.Carrinho
	_ = json.NewDecoder(r.Body).Decode(&carrinhoAtualizado)

	carrinho, ok := carrinhos[id]
	if !ok {
			http.Error(w, "Carrinho não encontrado", http.StatusNotFound)
			return
	}
	carrinho.UserId = carrinhoAtualizado.UserId
	carrinho.ValorTotal = carrinhoAtualizado.ValorTotal
	carrinho.InfosProduto = carrinhoAtualizado.InfosProduto
	carrinhos[id] = carrinho

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carrinho)
}

func deleteCarrinho(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, ok := carrinhos[id]
	if ok {
			delete(carrinhos, id)
	}

	if !ok {
			http.Error(w, "Carrinho não encontrado", http.StatusNotFound)
			return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createUsuario(w http.ResponseWriter, r *http.Request) {
	var novoUsuario model.Usuario
	_ = json.NewDecoder(r.Body).Decode(&novoUsuario)

	// Gerar um ID aleatório para o usuário
	novoUsuario.Id = strconv.Itoa(rand.Intn(1000000))

	// Hash da senha do usuário
	hashedSenha, err := hashSenha(novoUsuario.Senha)
	if err != nil {
			http.Error(w, "Erro ao hashear a senha", http.StatusInternalServerError)
			return
	}
	novoUsuario.Senha = hashedSenha

	// Armazenar o novo usuário
	mutex.Lock()
	usuarios[novoUsuario.Email] = novoUsuario
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novoUsuario)
}

func loginUsuario(w http.ResponseWriter, r *http.Request) {
	var credenciais struct {
			Email string `json:"email"`
			Senha string `json:"senha"`
	}
	_ = json.NewDecoder(r.Body).Decode(&credenciais)

	mutex.Lock()
	usuario, ok := usuarios[credenciais.Email]
	mutex.Unlock()

	if !ok || !verificaSenha(usuario.Senha, credenciais.Senha) {
			http.Error(w, "Email ou senha inválidos", http.StatusUnauthorized)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

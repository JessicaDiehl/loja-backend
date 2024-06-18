import (
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
	"sync"
)

// Função para hash da senha
func hashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

// Função para verificar a senha
func verificaSenha(hash, senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

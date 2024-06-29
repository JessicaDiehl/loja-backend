package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
	"github.com/google/uuid"
    "strconv"
    "time"
)

type Payment struct {
    ID        string  `json:"id"`
    Amount    float64 `json:"amount"`
    Status    string  `json:"status"`
    Timestamp string  `json:"timestamp"`
}

var payments []Payment

func createPagamento(w http.ResponseWriter, r *http.Request) {
    var payment Payment
    _ = json.NewDecoder(r.Body).Decode(&payment)
    payment.ID = uuid.New().String() // gera ID único
    payment.Status = "Processed"
    payment.Timestamp = time.Now().Format(time.RFC3339)
    payments = append(payments, payment)
    json.NewEncoder(w).Encode(payment)
}

func getPagamento(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    for _, payment := range payments {
        if payment.ID == id {
            json.NewEncoder(w).Encode(payment)
            return
        }
    }

    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Pagamento não encontrado"})
}
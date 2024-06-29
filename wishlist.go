package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
	"github.com/google/uuid"
    "strconv"
)

type WishlistItem struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var wishlist []WishlistItem

func createItemWishlist(w http.ResponseWriter, r *http.Request) {
    var item WishlistItem
    _ = json.NewDecoder(r.Body).Decode(&item)
    item.ID = uuid.New().String() // gera ID único
    wishlist = append(wishlist, item)
    json.NewEncoder(w).Encode(item)
}

func getItemsWishlist(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(wishlist)
}

func deleteItemWishlist(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    for index, item := range wishlist {
        if item.ID == id {
            wishlist = append(wishlist[:index], wishlist[index+1:]...)
            break
        }
    }

    json.NewEncoder(w).Encode(wishlist)
}
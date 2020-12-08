package router

import (
	"go-user/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

// Router users
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/addresses", api.GetAddress).Methods("GET") // Get Addresses
	router.HandleFunc("/user_address", api.UserAddress).Methods("GET")

	http.Handle("/", router)

	return router
}

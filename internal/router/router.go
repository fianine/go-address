package router

import (
	"go-user/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

// Router address
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/addresses", api.GetAddress).Methods("GET")
	router.HandleFunc("/user_address", api.UserAddress).Methods("GET")
	router.HandleFunc("/add_user_address", api.AddUserAddress).Methods("POST")

	http.Handle("/", router)

	return router
}

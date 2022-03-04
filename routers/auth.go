package routers

import (
	"github.com/afaferz/social-application/controllers"
	"github.com/gorilla/mux"
)

//AddRouterEndpoints add the actual endpoints for api
func AddRouterAuth(r *mux.Router) {
	r.HandleFunc("/api/auth/login/", controllers.GetTokenUserPassword).Methods("GET")
	r.HandleFunc("/api/auth/create-user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/auth/{POST_ID}/", controllers.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/auth/{POST_ID}/comments/", controllers.AddComment).Methods("POST")
}

package routers

import (
	"github.com/afaferz/social-application/controllers"
	"github.com/gorilla/mux"
)

//AddRouterEndpoints add the actual endpoints for api
func AddRouterPosts(r *mux.Router) {
	r.HandleFunc("/api/posts/", controllers.AllPosts).Methods("GET")
	r.HandleFunc("/api/posts/", controllers.AddPost).Methods("POST")
	r.HandleFunc("/api/posts/{POST_ID}/", controllers.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{POST_ID}/comments/", controllers.AddComment).Methods("POST")
}

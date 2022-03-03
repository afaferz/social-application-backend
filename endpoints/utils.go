package endpoints

import (
	"github.com/afaferz/social-application/controllers/posts"
	"github.com/gorilla/mux"
)

//AddRouterEndpoints add the actual endpoints for api
func AddRouterEndpoints(r *mux.Router) *mux.Router {

	r.HandleFunc("/api/posts/", posts.AllPosts).Methods("GET")
	r.HandleFunc("/api/posts/", posts.AddPost).Methods("POST")
	r.HandleFunc("/api/posts/{POST_ID}/", posts.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/posts/{POST_ID}/comments/", posts.AddComment).Methods("POST")
	return r
}

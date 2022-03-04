package routers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	AddRouterPosts(r)
	return r
}

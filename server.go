package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afaferz/social-application/middlewares"
	"github.com/afaferz/social-application/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r = routers.NewRouter()
	fs := http.FileServer(http.Dir("./dist"))
	r.PathPrefix("/").Handler(fs)

	http.Handle("/", &CORSRouterDecorator{r})
	fmt.Println("Listening")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}

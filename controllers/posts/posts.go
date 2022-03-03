package posts

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/afaferz/social-application/models/posts"
)

var postsArray []posts.Post = make([]posts.Post, 0)
var index int = 1

func AllPosts(wr http.ResponseWriter, req *http.Request) {
	log.Println("/api/posts/ - GET")
	json.NewEncoder(wr).Encode(postsArray)
}

func AddPost(wr http.ResponseWriter, req *http.Request) {
	log.Println("/api/posts/ - POST")
	var actualPost posts.Post
	err := json.NewDecoder(req.Body).Decode(&actualPost)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	actualPost.Id = index
	index++
	actualPost.Date = time.Now()
	if actualPost.Comments == nil {
		actualPost.Comments = make([]posts.Comment, 0)
	}
	postsArray = append(postsArray, actualPost)
	json.NewEncoder(wr).Encode(postsArray)
}

func DeletePost(wr http.ResponseWriter, req *http.Request) {
	log.Println("/api/posts/{POST_ID} - DELETE")
	vars := mux.Vars(req)
	idQueryParams, ok := vars["POST_ID"]
	if !ok {
		http.Error(wr, "Cannot find id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idQueryParams)
	if err != nil {
		http.Error(wr, "Cannot convert the id value to string", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(postsArray); i++ {
		if postsArray[i].Id == id {
			postsArray[i] = postsArray[len(postsArray)-1]
			postsArray = postsArray[:len(postsArray)-1]
			wr.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(wr, "Cannot find the requested id", http.StatusNotFound)
}

func AddComment(wr http.ResponseWriter, req *http.Request) {
	log.Println("/api/posts/{POST_ID}/comments/ - POST")
	vars := mux.Vars(req)
	idQueryParams, ok := vars["POST_ID"]
	if !ok {
		http.Error(wr, "Cannot find ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idQueryParams)
	if err != nil {
		http.Error(wr, "Cannot convert the id value to string", http.StatusBadRequest)
		return
	}
	var actualComment posts.Comment
	err = json.NewDecoder(req.Body).Decode(&actualComment)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < len(postsArray); i++ {
		if postsArray[i].Id == id {
			var commMax int = 0
			for comm := 0; comm < len(postsArray[i].Comments); comm++ {
				if commMax < postsArray[i].Comments[comm].Id {
					commMax = postsArray[i].Comments[comm].Id
				}
			}
			actualComment.Id = commMax + 1
			actualComment.Date = time.Now()
			postsArray[i].Comments = append(postsArray[i].Comments, actualComment)
			json.NewEncoder(wr).Encode(postsArray[i])
			return
		}
	}
	http.Error(wr, "Cannot find a post with the selected id", http.StatusNotFound)
}

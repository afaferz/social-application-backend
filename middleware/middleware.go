package middleware

import (
	"net/http"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	var SetApplicationJSON = func(wr http.ResponseWriter, req *http.Request) {
		wr.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// wr.WriteHeader(http.StatusOK)
		next.ServeHTTP(wr, req)
	}
	return http.HandlerFunc(SetApplicationJSON)
}

// func SendJSONResponse(w http.ResponseWriter, data interface{}) {
// 	body, err := json.Marshal(data)
// 	if err != nil {
// 		log.Printf("Failed to encode a JSON response: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	_, err = w.Write(body)
// 	if err != nil {
// 		log.Printf("Failed to write the response body: %v", err)
// 		return
// 	}
// }

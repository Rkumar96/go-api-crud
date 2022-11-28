package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// UsersRouter handle th users route
func UsersRouter(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
	path = strings.TrimPrefix(path, "/users/")
	fmt.Println(path)
	// if path != (bsontype.Type).String {
	// 	postError(w, http.StatusNotFound)
	// 	return
	// }

	// id := bson.ObjectIDHex(path)

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}

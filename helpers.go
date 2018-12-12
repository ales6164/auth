package auth

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func UserKeyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vars["key"] = context.Get(r, "userKey").(string)
		h.ServeHTTP(w, mux.SetURLVars(r, vars))
	})
}

func CollectionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collectionKey := mux.Vars(r)["collection"]

		// todo: check if has permission to edit
		// how is with permissions for editing collection kinds???
		//

		//h.ServeHTTP(w, mux.SetURLVars(r, vars))
	})
}

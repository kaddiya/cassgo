package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
)

func ListTodo(app *config.AppContainer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello!"))
	})
}

func InitRouter(app *config.AppContainer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todo/", ListTodo(app)).Methods("GET")
	return r

}

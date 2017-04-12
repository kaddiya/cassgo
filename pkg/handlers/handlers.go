package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
)

func ListTodo(app *config.AppContainer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := app.TodoService.GetTodo()
		fmt.Println(result)
		w.WriteHeader(200)
		w.Write([]byte("hello!"))
	})
}

func InitRouter(app *config.AppContainer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Todo/", ListTodo(app)).Methods("GET")
	return r

}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
)

func ListFoo(app *config.AppContainer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := app.FooServiceImpl.GetFoo()
		fmt.Println(result)
		w.Write([]byte("hello!"))
	})
}

func CreateFoo(app *config.AppContainer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.FooServiceImpl.CreateFoo()
		w.Write([]byte("A foo sucessfully created!"))
	})
}

func InitRouter(app *config.AppContainer) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/foo/", ListFoo(app)).Methods("GET")
	r.HandleFunc("/foo/", CreateFoo(app)).Methods("POST")
	return r

}

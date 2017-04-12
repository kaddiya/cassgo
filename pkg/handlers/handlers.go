package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// List will list all foos
func (app *AppContainer) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(app.AppName))
}

// Create will create a foo
func (app *AppContainer) Create(w http.ResponseWriter, r *http.Request) {

}

//InitRouter initialised the router
func (app *AppContainer) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foo/", app.List).Methods("GET")
	r.HandleFunc("/foo/", app.Create).Methods("POST")
	return r
}

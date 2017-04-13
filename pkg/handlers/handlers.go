package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
	"github.com/kaddiya/cassgo/pkg/models"
)

func ListTodo(app *config.AppContainer) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iter := app.Session.Query("select id,isdone,title from todo").Iter()
		var id int
		var title string
		var done bool
		var res []*models.Todo

		for iter.Scan(&id, &done, &title) {
			fmt.Printf("\n %d-%s-%t", id, title, done)
			t := &models.Todo{Title: title, IsDone: done}
			res = append(res, t)
		}
		if err := iter.Close(); err != nil {
			panic(err)
		}
		out, _ := json.Marshal(res)
		w.WriteHeader(200)
		w.Write([]byte(string(out)))
	})
}

func InitRouter(app *config.AppContainer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Todo/", ListTodo(app)).Methods("GET")
	return r

}

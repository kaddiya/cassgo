package main

import (
	"net/http"

	"github.com/kaddiya/cassgo/pkg/config"
	"github.com/kaddiya/cassgo/pkg/handlers"
	"github.com/kaddiya/cassgo/pkg/services"
)

func main() {
	a := &config.AppContainer{AppName: "CassGo", TodoService: &services.TodoServiceImpl{}}
	http.ListenAndServe(":3001", handlers.InitRouter(a))
}

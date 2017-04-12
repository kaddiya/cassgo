package main

import (
	"net/http"

	"github.com/kaddiya/cassgo/pkg/handlers"
)

func main() {
	a := &handlers.AppContainer{AppName: "CassGo"}
	http.ListenAndServe(":3001", a.InitRouter())
}

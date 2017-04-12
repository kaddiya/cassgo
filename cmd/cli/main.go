package main

import (
	"net/http"

	"github.com/kaddiya/cassgo/pkg/cassandra"
	"github.com/kaddiya/cassgo/pkg/config"
	"github.com/kaddiya/cassgo/pkg/handlers"
)

func main() {
	session, err := cassandra.InitSession()
	if err != nil {
		panic("cant establish connection to the cluster")
	}
	a := &config.AppContainer{AppName: "CassGo", Session: session}
	http.ListenAndServe(":3001", handlers.InitRouter(a))
}

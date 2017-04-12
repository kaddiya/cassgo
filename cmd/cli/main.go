package main

import (
	"fmt"

	"github.com/kaddiya/cassgo/pkg/container"
)

func main() {
	a := &container.AppContainer{AppName: "CassGo"}
	fmt.Printf("Welcome to %s\n", a.AppName)
}

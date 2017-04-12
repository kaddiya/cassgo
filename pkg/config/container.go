package config

import "github.com/gocql/gocql"

// AppContainer holds the dependencies that are required
type AppContainer struct {
	AppName string
	Session *gocql.Session
}

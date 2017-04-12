package cass

import (
	"github.com/gocql/gocql"
)

func InitGoCqlSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

type GoCQLSessionProvider interface {
}

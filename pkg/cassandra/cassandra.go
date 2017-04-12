package cassandra

import "github.com/gocql/gocql"

// InitSession will initialise the seesion
func InitSession(hosts ...string) (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

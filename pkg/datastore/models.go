package datastore

type Datastore interface {
	GetAllTodos([]*Todo, error)
}

type Todo struct {
}

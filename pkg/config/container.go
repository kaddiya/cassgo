package config

// AppContainer holds the dependencies that are required
type AppContainer struct {
	AppName    string
	FooService FooService
}

type FooService interface {
	GetFoo() []string
}

package services

type FooServiceImpl struct {
}

func (f *FooServiceImpl) GetFoo() []string {
	return []string{"1", "2", "3"}
}

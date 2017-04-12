package services

type TodoServiceImpl struct {
}

func (f *TodoServiceImpl) GetTodo() []string {
	return []string{"1", "2", "3"}
}

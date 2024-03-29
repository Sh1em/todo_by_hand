package service

import (
	todo "todo_app"
	"todo_app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetALL(userId int) ([]todo.TodoList, error) {
	return s.repo.GetALL(userId)
}

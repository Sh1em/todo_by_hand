package service

import (
	todo "todo_app"
	"todo_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}

type Service interface {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

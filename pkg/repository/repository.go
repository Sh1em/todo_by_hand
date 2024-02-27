package repository

import (
	todo "todo_app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}

type Repository interface {
	Authorization
	//	TodoList
	//	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

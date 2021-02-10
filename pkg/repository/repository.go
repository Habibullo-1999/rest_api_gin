package repository

import (
	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/jmoiron/sqlx"

)

type Authorization interface {
	CreateUser(user rest_api_gin.User) (int, error)
	GetUser(username, password string) (rest_api_gin.User, error)	
}

type NoteList interface {
}

type NoteItem interface {
}

type Repository struct {
	Authorization
	NoteList
	NoteItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		
	}
}
